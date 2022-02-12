package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"path/filepath"

	//"text/template"
	"github.com/lucasb-eyer/go-colorful"
)

/* Named colors taken from colorgorial implementation at:
http://vrl.cs.brown.edu/color

With all sliders maxed out.
Cite:
@article{gramazio-2017-ccd,
  author={Gramazio, Connor C. and Laidlaw, David H. and Schloss, Karen B.},
  journal={IEEE Transactions on Visualization and Computer Graphics},
  title={Colorgorical: creating discriminable and preferable color palettes for information visualization},
  year={2017}
}
*/

func main() {
	log.SetFlags(0)
	baseColorsHex := map[string]string{
		"":       "#EDD1B0",
		"orange": "#EDDD6E",
		"yellow": "#F8FD89",
	}
	fgHex := "#000000"

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("usage: mktheme <output directory>")
	}
	outputDir := os.Args[1]

	scheme := map[string]string{
		"problem":  "#c31a31",
		"notice":   "#d47a37",
		"affirm":   "#769d31",
		"meta":     "#6e5fae",
		"modified": "#b76f7c",
		"added":    "#1a8298",
		"removed":  "#cccccc",
	}

	for baseName, bgHex := range baseColorsHex {
		scheme["name"] = baseName
		bg, err := colorful.Hex(bgHex)
		if err != nil {
			log.Fatalf("%q is not a valid hex color", bgHex)
		}
		fg, err := colorful.Hex(fgHex)
		if err != nil {
			log.Fatalf("%q is not a valid hex color", fgHex)
		}
		bgl, bga, bgb := bg.Lab()
		fgl, fga, fgb := fg.Lab()
		colorFuncs := template.FuncMap{
			"bg": func(i int) string {
				return colorful.Lab(bgl-float64(i)*0.06, bga, bgb).Hex()
			},
			"fg": func(i int) string {
				return colorful.Lab(fgl+float64(i)*0.06, fga, fgb).Hex()
			},
			"alpha": func(pct int, c string) string {
				// add alpha in hex from range of 0-100%
				return fmt.Sprintf("%s%02x", c, int(math.Round(float64(pct)*255.0/100.0)))
			},
		}
		themeName := `clarion-color-theme`
		if baseName != "" {
			themeName += `-` + baseName
		}
		outPath := filepath.Join(outputDir, themeName+`.json`)
		outFile, err := os.Create(outPath)
		if err != nil {
			log.Fatalf("unable to create output file %q: %v", outPath, err)
		}
		tmpl, err := template.New("").Funcs(colorFuncs).ParseFiles("clarion-color-theme.json.tmpl")
		if err != nil {
			log.Fatalf("template parse error: %v", err)
		}
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json.tmpl", scheme); err != nil {
			log.Fatalf("template execution error: %v", err)
		}
	}

	// log.Printf("%#v\n", scheme)
}
