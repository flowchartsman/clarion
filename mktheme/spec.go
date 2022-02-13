package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type spec struct {
	themeBases    []string
	fgColor       string
	baseColors    map[string]string
	conceptColors map[string]string
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
		baseColors:    map[string]string{},
		conceptColors: map[string]string{},
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
					return nil, fmt.Errorf("couldn't parse DeltaETarget as float: %s", err)
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
					spec.baseColors[target] = fields[2]
				case "Foreground Colors":
					spec.fgColor = fields[2]
				case "Conceptual Colors":
					spec.conceptColors[target] = fields[2]
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %s", err)
	}
	return spec, nil
}
