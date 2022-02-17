package main

import (
	"fmt"
	"html/template"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// CIELAB ΔE* is the latest iteration of the CIE's color distance function, and
// is intended to meaure the perceived distance between two colors, where a
// value of 1.0 represents a "just noticeable difference".
// ref:  https://en.wikipedia.org/wiki/Color_difference#CIEDE2000

type renderTable struct {
	bgColors,
	bgColorsTerm map[string][]string
	fgColors,
	fgColorsTerm []string
	conceptColors,
	conceptColorsTerm map[string]string
}

func newRenderTable() renderTable {
	return renderTable{
		bgColors:          map[string][]string{},
		bgColorsTerm:      map[string][]string{},
		fgColors:          []string{},
		fgColorsTerm:      []string{},
		conceptColors:     map[string]string{},
		conceptColorsTerm: map[string]string{},
	}
}

func buildThemes(specPath string, outputPath string) error {
	spec, err := loadSpec(specPath)
	if err != nil {
		return fmt.Errorf("error loading specification: %s", err)
	}

	renderTable := newRenderTable()

	for baseColorName, baseColor := range spec.baseColors {
		hexVariations, termVariations, err := generateVariations(baseColor, spec.variations, spec.ΔETarget, spec.Lstep, both)
		if err != nil {
			return fmt.Errorf("error generating variations for base color %s: %s", baseColorName, err)
		}
		renderTable.bgColors[baseColorName] = hexVariations
		renderTable.bgColorsTerm[baseColorName] = termVariations
	}

	{
		fgVariationsHex, fgVariationsTerm, err := generateVariations(spec.fgColor, spec.variationsFG, spec.ΔETargetFG, spec.Lstep, lighter)
		if err != nil {
			return fmt.Errorf("error generating variations for foreground color: %s", err)
		}
		renderTable.fgColors = fgVariationsHex
		renderTable.fgColorsTerm = fgVariationsTerm
	}

	for conceptColorName, conceptColorStr := range spec.conceptColors {
		conceptColor, _ := colorful.Hex(conceptColorStr)
		renderTable.conceptColors[conceptColorName] = conceptColor.Hex()
		renderTable.conceptColorsTerm[conceptColorName] = toTerminal(conceptColor)
	}

	for i, baseColor := range spec.themeBases {
		themeFileSuffix := "-" + strings.ToLower(baseColor)
		if i == 0 {
			themeFileSuffix = ""
		}
		// Create Template functions
		colorFuncs := template.FuncMap{
			"themeName": func() string {
				return "Clarion " + baseColor
			},
			"bg": func(offset int) string {
				center := len(renderTable.bgColors[baseColor]) / 2
				idx := center + offset
				if idx < 0 || idx >= len(renderTable.bgColors[baseColor]) {
					themeLogErr("bg idx out of range: offset: %d idx: %d", offset, idx)
					return errColor
				}
				return renderTable.bgColors[baseColor][idx]
			},
			"bg256": func(offset int) string {
				center := len(renderTable.bgColorsTerm[baseColor]) / 2
				idx := center + offset
				if idx < 0 || idx >= len(renderTable.bgColorsTerm[baseColor]) {
					themeLogErr("bg256 idx out of range: offset: %d idx: %d", offset, idx)
					return errColor256
				}
				return renderTable.bgColorsTerm[baseColor][idx]
			},
			"fg": func(offset int) string {
				if offset < 0 || offset >= len(renderTable.fgColors) {
					themeLogErr("fg offset out of range: %d", offset)
					return errColor
				}
				return renderTable.fgColors[offset]
			},
			"fg256": func(offset int) string {
				if offset < 0 || offset >= len(renderTable.fgColorsTerm) {
					themeLogErr("fg256 offset out of range: %d", offset)
					return errColor
				}
				return renderTable.fgColorsTerm[offset]
			},
			"alpha": func(pct int, c string) string {
				// add alpha in hex from range of 0-100%
				return fmt.Sprintf("%s%02x", c, int(math.Round(float64(pct)*255.0/100.0)))
			},
		}
		for conceptColorName := range spec.conceptColors {
			colorname := conceptColorName
			colorFuncs[colorname] = func() string {
				return renderTable.conceptColors[colorname]
			}
			colorFuncs[colorname+"256"] = func() string {
				return renderTable.conceptColorsTerm[colorname]
			}
		}
		themeFilename := fmt.Sprintf("clarion-color-theme%s.json", themeFileSuffix)
		outPath := filepath.Join(outputPath, themeFilename)
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("unable to create output file %q: %v", outPath, err)
		}
		defer outFile.Close()
		tmpl, err := template.New("").Funcs(colorFuncs).ParseFiles("template/clarion-color-theme.json")
		if err != nil {
			return fmt.Errorf("template parse error: %v", err)
		}
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json", nil); err != nil {
			return fmt.Errorf("template execution error: %v", err)
		}
	}
	return nil
}
