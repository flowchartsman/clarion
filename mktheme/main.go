package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/lucasb-eyer/go-colorful"
	//"text/template"
)

// CIELAB ΔE* is the latest iteration of the CIE's color distance function, and
// is intended to meaure the perceived distance between two colors, where a
// value of 1.0 represents a "just noticeable difference".
// ref:  https://en.wikipedia.org/wiki/Color_difference#CIEDE2000

// ΔETarget is the target perceptual distance we want to achieve for each
// permutation of a color. The color's L value is increased by LStep until this is achieved.
var ΔETarget = 3.0

const LStep = 0.1
const NumVariations = 6

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

func generateVariations(spec *spec, baseColorStr string, numVariations int, lighter bool) (hexVariations []string, termVariations []string, err error) {
	baseColor, _ := colorful.Hex(baseColorStr)
	hexVariations = make([]string, numVariations)
	termVariations = make([]string, numVariations)
	lastVariation := baseColor
	for i := 0; i < NumVariations; i++ {
		variation := lastVariation
		if i == 0 {
			variation = baseColor
		} else {
			var distance float64
			for distance < spec.ΔETarget {
				l, a, b := variation.Lab()
				if lighter {
					l += spec.lStep
				} else {
					l -= spec.lStep
				}
				switch {
				case l >= 100:
					return nil, nil, fmt.Errorf("L overflow for variant %d", i)
				case l <= 0:
					return nil, nil, fmt.Errorf("L underflow for variant %d", i)
				}
				variation = colorful.Lab(l, a, b)
				distance = lastVariation.DistanceCIEDE2000(variation)
			}
		}
		variation = variation.Clamped()
		hexVariations[i] = variation.Hex()
		termVariations[i] = toTerminal(variation)
		lastVariation = variation
	}
	return hexVariations, termVariations, nil
}

func main() {
	log.SetFlags(0)
	if len(os.Args) != 3 {
		log.Fatalf("usage: mktheme <spec markdown file> <output directory>")
	}
	spec, err := loadSpec(os.Args[1])
	if err != nil {
		log.Fatalf("error loading specification: %s", err)
	}
	outputDir := os.Args[2]

	log.Println("generating color permutations")
	renderTable := newRenderTable()

	for baseColorName, baseColor := range spec.baseColors {
		hexVariations, termVariations, err := generateVariations(spec, baseColor, NumVariations, true)
		if err != nil {
			log.Fatalf("error generating variations for base color %s: %s", baseColorName, err)
		}
		renderTable.bgColors[baseColorName] = hexVariations
		renderTable.bgColorsTerm[baseColorName] = termVariations
	}

	{
		fgVariationsHex, fgVariationsTerm, err := generateVariations(spec, spec.fgColor, NumVariations, true)
		if err != nil {
			log.Fatalf("error generating variations for foreground color: %s", err)
		}
		renderTable.fgColors = fgVariationsHex
		renderTable.fgColorsTerm = fgVariationsTerm
	}

	for conceptColorName, conceptColorStr := range spec.conceptColors {
		conceptColor, _ := colorful.Hex(conceptColorStr)
		renderTable.conceptColors[conceptColorName] = conceptColor.Hex()
		renderTable.conceptColorsTerm[conceptColorName] = toTerminal(conceptColor)
	}

	log.Printf("%#v\n", renderTable)

	for i, baseColor := range spec.themeBases {
		themeSuffix := baseColor
		if i == 0 {
			themeSuffix = ""
		}
		// Create Template functions
		colorFuncs := template.FuncMap{
			"themeName": func() string {
				return themeSuffix
			},
			"bg": func(i int) string {
				return renderTable.bgColors[baseColor][i]
			},
			"bg256": func(i int) string {
				return renderTable.bgColorsTerm[baseColor][i]
			},
			"fg": func(i int) string {
				return renderTable.fgColors[i]
			},
			"fg256": func(i int) string {
				return renderTable.fgColorsTerm[i]
			},
			"alpha": func(pct int, c string) string {
				// add alpha in hex from range of 0-100%
				return fmt.Sprintf("%s%02x", c, int(math.Round(float64(pct)*255.0/100.0)))
			},
		}
		for conceptColorName := range spec.conceptColors {
			colorFuncs[conceptColorName] = func() string {
				return renderTable.conceptColors[conceptColorName]
			}
			colorFuncs[conceptColorName+"256"] = func() string {
				return renderTable.conceptColorsTerm[conceptColorName]
			}
		}
		themeName := `clarion-color-theme` + themeSuffix
		outPath := filepath.Join(outputDir, themeName+`.json`)
		outFile, err := os.Create(outPath)
		if err != nil {
			log.Fatalf("unable to create output file %q: %v", outPath, err)
		}
		tmpl, err := template.New("").Funcs(colorFuncs).ParseFiles("clarion-color-theme.json.tmpl")
		if err != nil {
			log.Fatalf("template parse error: %v", err)
		}
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json.tmpl", nil); err != nil {
			log.Fatalf("template execution error: %v", err)
		}
	}
	// log.Printf("%#v\n", scheme)
}
