package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// todo:
// allow for .BG(1) and .BG(1).Problem
// this can be done by making the colors methods of the BG, such that *BGType
// has .String()string method and methods for colors This would allow a really
// nice
// {{ with .BG 1}}bg = {{.}} fg = {{.Problem.Txt}}{{end}}
// or {{ with .Level 1}}...{{end}}

type themeColor struct {
	c colorful.Color
}

func (tc *themeColor) Hex() string {
	return tc.c.Clamped().Hex()
}

func (tc *themeColor) HexT() string {
	return tc.Hex()[1:]
}

// X * Y/100. x percent of y
func (tc *themeColor) HexAlpha(transPct float64) string {
	return fmt.Sprintf("%s%02x", tc.Hex(), int(math.Round(transPct*255.0/100.0)))
}

func (tc *themeColor) Term256() string {
	if tc.c.IsValid() {
		return toTerminal(tc.c.Clamped())
	}
	return toTerminal(tc.c)
}

type Background struct {
	*themeColor
	Fg      *themeColor
	FgLight *themeColor
	FgUI    *themeColor
	cColors map[string]*colorElement
	ansi    map[string]*themeColor
	mytheme *Theme
	myIdx   int
}

func (b *Background) Darker() *Background {
	dIdx := b.myIdx - 1
	if dIdx < 0 {
		themeLogErr("Darker(): idx is <0: %d", dIdx)
		return ErrColor
	}
	return b.mytheme.backgrounds[dIdx]
}

func (b *Background) Lighter() *Background {
	lIdx := b.myIdx + 1
	if lIdx >= len(b.mytheme.backgrounds) {
		themeLogErr("Lighter(): idx is out of bounds: %d", lIdx)
		return ErrColor
	}
	return b.mytheme.backgrounds[lIdx]
}

// new methods here, as needed
func (b *Background) Problem() *colorElement {
	return b.cColors["problem"]
}

func (b *Background) Notice() *colorElement {
	return b.cColors["notice"]
}

func (b *Background) New() *colorElement {
	return b.cColors["new"]
}

func (b *Background) Modified() *colorElement {
	return b.cColors["modified"]
}

func (b *Background) Excluded() *colorElement {
	return b.cColors["excluded"]
}

func (b *Background) AnsiBlack() *themeColor {
	return b.ansi["Black"]
}

func (b *Background) AnsiBrBlack() *themeColor {
	return b.ansi["BrBlack"]
}

func (b *Background) AnsiRed() *themeColor {
	return b.ansi["Red"]
}

func (b *Background) AnsiBrRed() *themeColor {
	return b.ansi["BrRed"]
}

func (b *Background) AnsiGreen() *themeColor {
	return b.ansi["Green"]
}

func (b *Background) AnsiBrGreen() *themeColor {
	return b.ansi["BrGreen"]
}

func (b *Background) AnsiBlue() *themeColor {
	return b.ansi["Blue"]
}

func (b *Background) AnsiBrBlue() *themeColor {
	return b.ansi["BrBlue"]
}

func (b *Background) AnsiYellow() *themeColor {
	return b.ansi["Yellow"]
}

func (b *Background) AnsiBrYellow() *themeColor {
	return b.ansi["BrYellow"]
}

func (b *Background) AnsiMagenta() *themeColor {
	return b.ansi["Magenta"]
}

func (b *Background) AnsiBrMagenta() *themeColor {
	return b.ansi["BrMagenta"]
}

func (b *Background) AnsiCyan() *themeColor {
	return b.ansi["Cyan"]
}

func (b *Background) AnsiBrCyan() *themeColor {
	return b.ansi["BrCyan"]
}

func (b *Background) AnsiWhite() *themeColor {
	return b.ansi["White"]
}

func (b *Background) AnsiBrWhite() *themeColor {
	return b.ansi["BrWhite"]
}

type colorElement struct {
	Unmodified *themeColor
	Txt        *themeColor
	// enhanced  colorful.Color // rename text (but as variant)
	Gfx *themeColor
}

type Theme struct {
	// Variant is the name for this theme variant, based on its primary
	// seed (background) color
	Variant string
	// ThemeName is the full name for this variant for display
	ThemeName string
	// Version is the overall version of the theme as tagged in Git (or the hash)
	Version     string
	baseIdx     int
	backgrounds []*Background
}

func (t *Theme) Base() *Background {
	return t.backgrounds[t.baseIdx]
}

// BG represents a background variant indexed by positive integer (lighter) or
// negative integer (darker). Invalid indices will return ErrColor and a warning
// Bg(0) is the same a Base()
func (t *Theme) Bg(idx int) *Background {
	if idx == 0 {
		return t.backgrounds[t.baseIdx]
	}
	idx = t.baseIdx + idx
	if idx < 0 || idx >= len(t.backgrounds) {
		return ErrColor
	}
	return t.backgrounds[idx]
}

// func (v *variant) Problem() colorful

// TODO generate AA and AAA variants as "<FullName>" and "<FullName> - High
// Contrast", respectively
func generateVariants(config *MkthemeConfig, spec *spec) ([]*Theme, error) {
	variants := make([]*Theme, 0, len(spec.baseColors))
	for baseColorName, baseColor := range spec.baseColors {
		variantName := strings.ToLower(baseColorName)

		vErr := func(s string, v ...interface{}) error {
			return fmt.Errorf("variant %s - %s", variantName, fmt.Sprintf(s, v...))
		}

		// generate background color variations
		darks, err := generateBrightnessVariations(baseColor, spec.variations, spec.ΔETarget, spec.Lstep, darker)
		if err != nil {
			return nil, err
		}
		lights, err := generateBrightnessVariations(baseColor, spec.variations, spec.ΔETarget, spec.Lstep, lighter)
		if err != nil {
			return nil, err
		}
		// sanity check the foreground against the darkest background color
		lowestFGContrast := contrast(spec.fgColor, darks[len(darks)-1])
		if lowestFGContrast < 4.5 {
			return nil, vErr("contrast ratio of foreground color (%s) to darkest background variant (%s) is too low %f < 4.5", spec.fgColor.Hex(), darks[len(darks)-1].Hex(), lowestFGContrast)
		}

		// combine them to make the loop easier
		allBgs := make([]*Background, 0, len(darks)+len(lights)+1)
		for di := len(darks) - 1; di >= 0; di-- {
			allBgs = append(allBgs, &Background{themeColor: &themeColor{darks[di]}})
		}
		allBgs = append(allBgs, &Background{themeColor: &themeColor{baseColor}})

		for li := range lights {
			allBgs = append(allBgs, &Background{themeColor: &themeColor{lights[li]}})
		}

		// generate the theme-specific concept colors, for each background
		// TODO: extra loop for High contrast
		// TODO: fail if ΔE is too low from FG or others in the same BG cohort
		// (AA level only, warn for AAA)
		for bi, bg := range allBgs {
			bgColor := bg.c

			// generate colors at the upper two contrast levels for use as a
			// light foregrund text color and a UI element color.
			// TODO: consider color wheel transforms.
			fgUI, fgLight, _ := getContrastingColors(bgColor, bgColor)
			bg.FgUI = &themeColor{fgUI}
			bg.FgLight = &themeColor{fgLight}
			bg.Fg = &themeColor{spec.fgColor}

			// Concept Colors
			bg.cColors = map[string]*colorElement{}
			for ccName, ccColor := range spec.conceptColors {
				ccGfx, ccTxt, _ := getContrastingColors(bgColor, ccColor)
				bg.cColors[ccName] = &colorElement{
					Unmodified: &themeColor{ccColor},
					Txt:        &themeColor{ccTxt},
					Gfx:        &themeColor{ccGfx},
				}
			}

			// Get ANSI Colors
			ansiColors := map[string]*themeColor{
				"Black": {black},
			}
			brBlack, err := getNextBrightest(black, 0.02, spec.Lstep)
			if err != nil {
				return nil, vErr("generating brightness variations for ansiBlack: %v", err)
			}
			ansiColors["BrBlack"] = &themeColor{brBlack}
			// sanity check brightBlack against the background color
			brBlackContrast := contrast(brBlack, bgColor)
			if brBlackContrast < float64(4.5) {
				return nil, vErr("contrast ratio of ansiBrightBlack color (%s) to background (%s) is too low %f < 4.5", brBlack.Hex(), bgColor.Hex(), brBlackContrast)
			}
			for ansiColorName, ansiColorSpec := range spec.ansiColors {
				ansiColorBright, _, _ := getContrastingColors(bgColor, ansiColorSpec) // br color at Ui contrast level TODO: Fix
				ansiColors["Br"+ansiColorName] = &themeColor{ansiColorBright}
				// darken for normal ansi variant
				ansiColor, err := getNextDarkest(ansiColorBright, spec.ΔETarget, spec.Lstep)
				if err != nil {
					vErr("generating brightness variations for %s: %v", ansiColorName, err)
				}
				ansiColors[ansiColorName] = &themeColor{ansiColor}
			}
			bg.ansi = ansiColors
			bg.myIdx = bi
		}

		newTheme := &Theme{
			Variant:     variantName,
			ThemeName:   "Clarion " + baseColorName,
			Version:     config.themeVersion,
			baseIdx:     len(darks),
			backgrounds: allBgs,
		}
		for bi := range allBgs {
			allBgs[bi].mytheme = newTheme
		}
		variants = append(variants, newTheme)
	}
	return variants, nil
}

var (
	errColor = colorful.Hsl(300, 1, 0.5) // ugly ol' magenta
	black    = colorful.Hsl(0, 0, 0)
	ErrColor = &Background{
		themeColor: &themeColor{errColor},
		Fg:         &themeColor{errColor},
		FgLight:    &themeColor{errColor},
		FgUI:       &themeColor{errColor},
		cColors: map[string]*colorElement{
			"problem":  errElement,
			"notice":   errElement,
			"new":      errElement,
			"modified": errElement,
			"excluded": errElement,
		},
	}
	errElement = &colorElement{
		Unmodified: &themeColor{errColor},
		Txt:        &themeColor{black},
		Gfx:        &themeColor{black},
	}
)
