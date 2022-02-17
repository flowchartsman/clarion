package main

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

type adjustmentDirection int

const (
	_ adjustmentDirection = iota
	lighter
	darker
	both
)

func rev(s []string) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

const (
	errColor    = `#FF00FF`
	errColor256 = `201`
)

func generateVariations(baseColorStr string, variations int, ΔETarget float64, LStep float64, direction adjustmentDirection) (hexVariations []string, termVariations []string, err error) {
	if direction == both {
		variations /= 2
		hvl, tvl, err := generateVariations(baseColorStr, variations, ΔETarget, LStep, lighter)
		if err != nil {
			return nil, nil, err
		}
		hvd, tvd, err := generateVariations(baseColorStr, variations, ΔETarget, LStep, darker)
		if err != nil {
			return nil, nil, err
		}
		rev(hvd)
		rev(tvd)
		return append(hvd, hvl[1:]...), append(tvd, tvl[1:]...), nil
	}
	baseColor, _ := colorful.Hex(baseColorStr)
	hexVariations = make([]string, variations)
	termVariations = make([]string, variations)
	lastVariation := baseColor
	for i := 0; i < variations; i++ {
		variation := lastVariation
		if i == 0 {
			variation = baseColor
		} else {
			var distance float64
			for distance < ΔETarget {
				l, a, b := variation.Lab()
				if direction == lighter {
					l += LStep
				} else {
					l -= LStep
				}
				switch {
				case l >= 100:
					return nil, nil, fmt.Errorf("overflow L for variant %d", i)
				case l <= 0:
					return nil, nil, fmt.Errorf("underflow L for variant %d", i)
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
