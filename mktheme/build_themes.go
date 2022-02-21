package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
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
}

func buildThemes(specPath string, outputPath string) error {
	spec, err := loadSpec(specPath)
	if err != nil {
		return fmt.Errorf("error loading specification: %s", err)
	}

	pkg := &ThemePkg{
		Version: themeVersion,
	}

	masterTable := map[string]colorTable{}

	for baseColorName, baseColorHex := range spec.baseColors {
		masterTable[baseColorName] = colorTable{
			conceptColors: make(map[string]colorLevels),
		}
		baseColor, err := colorful.Hex(baseColorHex)
		if err != nil {
			return fmt.Errorf("base color hex invaid: %s", err)
		}
		fgColor, err := colorful.Hex(spec.fgColor)
		if err != nil {
			return fmt.Errorf("foreground color hex invalid: %s", err)
		}
		backgroundVariations, err := generateBackgroundVariations(baseColor, spec.variations, spec.ΔETarget, spec.Lstep, darker)
		if err != nil {
			return err
		}

		// sanity check the foreground against the darkest background color
		lowestFGContrast := contrast(fgColor, backgroundVariations[len(backgroundVariations)-1])
		if lowestFGContrast < 4.5 {
			return fmt.Errorf("contrast ratio of foreground color (%s) to darkest background variant (%s) is too low %f<4.5", fgColor.Hex(), backgroundVariations[len(backgroundVariations)-1].Hex(), lowestFGContrast)
		}

		// get the darkest background variation for calculating minimum contrast
		// for concept colors and ui elements
		darkestBackground := backgroundVariations[len(backgroundVariations)-1]

		cColors := make(map[string]colorLevels)
		for conceptColorName, conceptColorStr := range spec.conceptColors {
			conceptColor, err := colorful.Hex(conceptColorStr)
			if err != nil {
				return fmt.Errorf("concept color %q hex is invalid: %s", conceptColorName, err)
			}
			// generate concept colors based on the darkest background variation
			ccUI, ccMin, ccEnhanced := getLevels(darkestBackground, conceptColor)
			cColors[conceptColorName] = colorLevels{
				min:      ccMin,
				enhanced: ccEnhanced,
				ui:       ccUI,
			}
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
			fg:            fgColor,
			lightFg:       lightFg,
			conceptColors: cColors,
		}
	}

	for i, baseColor := range spec.themeBases {
		themeFileSuffix := "-" + strings.ToLower(baseColor)
		if i == 0 {
			themeFileSuffix = ""
		}
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
		defer outFile.Close()
		tmpl, err := template.New("").Funcs(colorFuncs).ParseFiles("template/clarion-color-theme.json")
		if err != nil {
			return fmt.Errorf("template parse error: %v", err)
		}
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json", nil); err != nil {
			return fmt.Errorf("template execution error: %v", err)
		}
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
	return nil
}
