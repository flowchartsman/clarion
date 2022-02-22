package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type spec struct {
	themeBases    []string
	fgColor       colorful.Color
	baseColors    map[string]colorful.Color
	conceptColors map[string]colorful.Color
	ansiColors    map[string]colorful.Color
	ΔETarget      float64
	Lstep         float64
	variations    int
}

// primitive, expect-like scanner that only works because markdown is easy. If
// the spec becomes more complicated, a full parser will likely be required.
func loadSpec(specPath string) (*spec, error) {
	file, err := os.Open(specPath)
	if err != nil {
		return nil, err
	}
	spec := &spec{
		baseColors:    map[string]colorful.Color{},
		conceptColors: map[string]colorful.Color{},
		ansiColors:    map[string]colorful.Color{},
	}
	scanner := bufio.NewScanner(file)
	section := "none"
	target := ""
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case `##`:
			section = scanner.Text()[3:]
			continue
		case `###`:
			target = scanner.Text()[4:]
		case `*`:
			switch fields[1] {
			case "ΔETarget:":
				f, err := strconv.ParseFloat(fields[2], 64)
				if err != nil {
					return nil, fmt.Errorf("couldn't parse ΔETarget as float: %s", err)
				}
				spec.ΔETarget = f / 100 //apparently the colorful library is a
				//couple orders of magnitude smaller than the literature
			case "LStep:":
				f, err := strconv.ParseFloat(fields[2], 64)
				if err != nil {
					return nil, fmt.Errorf("couldn't parse LStep as float: %s", err)
				}
				spec.Lstep = f / 100
			case "Variations:":
				i, err := strconv.Atoi(fields[2])
				if err != nil {
					return nil, fmt.Errorf("couldn't parse Variations as int")
				}
				spec.variations = i
			case "Hex:":
				switch section {
				case "Background Colors":
					spec.themeBases = append(spec.themeBases, target)
					bgHex := fields[2]
					bgColor, err := colorful.Hex(bgHex)
					if err != nil {
						return nil, fmt.Errorf("invalid hex [%s] for background color [%s]", bgHex, target)
					}
					spec.baseColors[target] = bgColor
				case "Foreground Colors":
					fgHex := fields[2]
					fgColor, err := colorful.Hex(fgHex)
					if err != nil {
						return nil, fmt.Errorf("invalid hex [%s] for foreground color", fgHex)
					}
					spec.fgColor = fgColor
				case "Conceptual Colors":
					ccHex := fields[2]
					cColor, err := colorful.Hex(ccHex)
					if err != nil {
						return nil, fmt.Errorf("invalid hex [%s] for concept color [%s]", ccHex, target)
					}
					spec.conceptColors[strings.ToLower(target)] = cColor
				case "Terminal Colors":
					tHex := fields[2]
					tColor, err := colorful.Hex(tHex)
					if err != nil {
						return nil, fmt.Errorf("invalid hex [%s] for terminal color [%s]", tHex, target)
					}
					spec.ansiColors[target] = tColor
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %s", err)
	}
	return spec, nil
}
