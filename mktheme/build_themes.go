package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

const themeVersion = `0.0.1`

type ThemePkg struct {
	Version       string
	ThemeContribs []ThemeContrib
}

type ThemeContrib struct {
	Label string
	File  string
}

type colorLevels struct {
	min      colorful.Color
	enhanced colorful.Color
	ui       colorful.Color
}

type colorTable struct {
	ui            colorful.Color
	bg            []colorful.Color
	fg            colorful.Color
	lightFg       colorful.Color
	conceptColors map[string]colorLevels
	ansiColors    map[string]colorful.Color
}

func buildThemes(spec *spec, outputPath string) error {
	pkg := &ThemePkg{
		Version: themeVersion,
	}

	masterTable := map[string]colorTable{}

	for baseColorName, baseColor := range spec.baseColors {
		// generate background color variations
		backgroundVariations, err := generateBrightnessVariations(baseColor, spec.variations, spec.ΔETarget, spec.Lstep, darker)
		if err != nil {
			return err
		}

		// sanity check the foreground against the darkest background color
		lowestFGContrast := contrast(spec.fgColor, backgroundVariations[len(backgroundVariations)-1])
		if lowestFGContrast < 4.5 {
			return fmt.Errorf("contrast ratio of foreground color (%s) to darkest background variant (%s) is too low %f < 4.5", spec.fgColor.Hex(), backgroundVariations[len(backgroundVariations)-1].Hex(), lowestFGContrast)
		}

		// get the darkest background variation for calculating minimum contrast
		// for concept colors and ui elements
		darkestBackground := backgroundVariations[len(backgroundVariations)-1]

		// generate the theme-specific concept colors
		cColors := make(map[string]colorLevels)
		for conceptColorName, conceptColor := range spec.conceptColors {
			// generate concept colors based on the darkest background variation
			ccUI, ccMin, ccEnhanced := getLevels(darkestBackground, conceptColor)
			cColors[conceptColorName] = colorLevels{
				min:      ccMin,
				enhanced: ccEnhanced,
				ui:       ccUI,
			}
		}

		// generate the theme-specific ansi colors
		black := colorful.Color{R: 0, G: 0, B: 0}
		ansiColors := map[string]colorful.Color{
			"ansiBlack": black,
		}
		blackVariations, err := generateBrightnessVariations(black, 2, 0.02, spec.Lstep, lighter)
		if err != nil {
			return err
		}
		brightBlack := blackVariations[1]
		if !brightBlack.IsValid() {
			brightBlack = brightBlack.Clamped()
		}
		ansiColors["ansiBrightBlack"] = brightBlack

		// sanity check brightBlack against the darkest background color
		brightBlackContrast := contrast(brightBlack, backgroundVariations[len(backgroundVariations)-1])
		if brightBlackContrast < float64(4.5) {
			return fmt.Errorf("contrast ratio of ansiBrightBlack color (%s) to darkest background variant (%s) is too low %f < 4.5", brightBlack.Hex(), backgroundVariations[len(backgroundVariations)-1].Hex(), brightBlackContrast)
		}

		for ansiColorName, ansiColor := range spec.ansiColors {
			themeAnsi, _, _ := getLevels(darkestBackground, ansiColor)
			ansiColors["ansiBright"+ansiColorName] = themeAnsi
			ansiVariations, err := generateBrightnessVariations(themeAnsi, 2, spec.ΔETarget, spec.Lstep, darker)
			if err != nil {
				return err
			}
			ansiDark := ansiVariations[1]
			if !ansiDark.IsValid() {
				ansiDark = ansiDark.Clamped()
			}
			ansiColors["ansi"+ansiColorName] = ansiDark
		}

		//generate a color for ui elements based on ui contrast level for
		//darkest background variation
		uiElementColor, _, _ := getLevels(darkestBackground, darkestBackground)

		// generate the lightest possible fg color we can use in the main editor
		// area
		_, lightFg, _ := getLevels(baseColor, baseColor)

		masterTable[baseColorName] = colorTable{
			ui:            uiElementColor,
			bg:            backgroundVariations,
			fg:            spec.fgColor,
			lightFg:       lightFg,
			conceptColors: cColors,
			ansiColors:    ansiColors,
		}
	}

	for _, baseColor := range spec.themeBases {
		themeFileSuffix := "-" + strings.ToLower(baseColor)
		ThemeName := "Clarion " + baseColor

		themeTable := masterTable[baseColor]
		// Create Template functions
		colorFuncs := template.FuncMap{
			"themeName": func() string {
				return ThemeName
			},
			"bg":       tmplBG(themeTable),
			"fg":       tmplFG(themeTable),
			"lightfg":  tmplLightFG(themeTable),
			"uifg":     tmplUIFG(themeTable),
			"hex":      tmpl2Hex,
			"hexalpha": tmpl2HexAlpha,
			"term":     tmpl2Term,
		}
		for conceptColorName := range themeTable.conceptColors {
			colorFuncs[conceptColorName] = tmplColor(conceptColorName, themeTable)
		}
		for ansiColorName := range themeTable.ansiColors {
			colorFuncs[ansiColorName] = tmplAnsiColor(ansiColorName, themeTable)
		}
		themeFilename := fmt.Sprintf("clarion-color-theme%s.json", themeFileSuffix)
		pkg.ThemeContribs = append(pkg.ThemeContribs, ThemeContrib{
			Label: ThemeName,
			File:  themeFilename,
		})
		outPath := filepath.Join(outputPath, "themes", themeFilename)
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("unable to create output file %q: %v", outPath, err)
		}
		tmpl, err := template.New("").Funcs(colorFuncs).ParseFiles("template/clarion-color-theme.json")
		if err != nil {
			return fmt.Errorf("template parse error: %v", err)
		}
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json", nil); err != nil {
			return fmt.Errorf("template execution error: %v", err)
		}
		outFile.Close()
	}
	pkgPath := filepath.Join(outputPath, "package.json")
	outPkg, err := os.Create(pkgPath)
	if err != nil {
		return fmt.Errorf("unable to create package output file %q: %v", pkgPath, err)
	}
	defer outPkg.Close()
	tmpl, err := template.New("").ParseFiles("template/package.json.tmpl")
	if err != nil {
		return fmt.Errorf("package template parse error: %v", err)
	}
	if err := tmpl.ExecuteTemplate(outPkg, "package.json.tmpl", pkg); err != nil {
		return fmt.Errorf("package template execution error: %v", err)
	}
	// generate readme
	readmePreviews := []map[string]string{}
	for _, base := range spec.themeBases {
		readmePreviews = append(readmePreviews, map[string]string{
			"themename":  "Clarion " + base,
			"screenshot": fmt.Sprintf("Clarion-%s.jpg", base),
		})
	}
	sort.Slice(readmePreviews, func(i, j int) bool {
		if readmePreviews[i]["themename"] == "Clarion White" {
			return true
		}
		if readmePreviews[j]["themename"] == "Clarion White" {
			return false
		}
		return readmePreviews[i]["themename"] < readmePreviews[j]["themename"]
	})
	tmpl, err = template.New("").ParseFiles(`template/README.md.tmpl`)
	if err != nil {
		return err
	}
	readmeout, err := os.Create(filepath.Join(outputPath, "README.md"))
	if err != nil {
		return err
	}
	defer readmeout.Close()
	if err := tmpl.ExecuteTemplate(readmeout, "README.md.tmpl", readmePreviews); err != nil {
		return err
	}
	return nil
}
