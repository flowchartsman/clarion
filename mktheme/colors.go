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
)

// func rev(s []string) {
// 	for i := len(s)/2 - 1; i >= 0; i-- {
// 		opp := len(s) - 1 - i
// 		s[i], s[opp] = s[opp], s[i]
// 	}
// }

func sRGBLuminance(color colorful.Color) float64 {
	r, g, b := color.LinearRgb()
	return 0.2126*r + 0.7152*g + 0.0722*b
}

func contrast(c1 colorful.Color, c2 colorful.Color) float64 {
	l1 := sRGBLuminance(c1)
	l2 := sRGBLuminance(c2)
	if l2 > l1 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05)
}

const lStep = 0.001

func getLevels(bg, fg colorful.Color) (ui, minimum, enhanced colorful.Color) {
	current := fg
	targetContrasts := [3]float64{3.0, 4.5, 7}
	outputs := [3]*colorful.Color{&ui, &minimum, &enhanced}
	levelStr := [3]string{"ui", "minimum", "maximum"}
	for i := 0; i < 3; i++ {
		for contrast(bg, current) < targetContrasts[i] {
			if current.R == 0 && current.G == 0 && current.B == 0 {
				themeLog("fg color has bottomed out bg[%s] fg[%s] variant[%s]", bg.Hex(), fg.Hex(), levelStr[i])
				for ii := i; ii < 3; ii++ {
					*outputs[ii] = current
				}
				return
			}
			h, s, l := current.HSLuv()
			l -= lStep
			current = colorful.HSLuv(h, s, l)
		}
		*outputs[i] = current
	}
	return
}

// CIELAB ΔE* is the latest iteration of the CIE's color distance function, and
// is intended to meaure the perceived distance between two colors, where a
// value of 1.0 represents a "just noticeable difference".
// ref:  https://en.wikipedia.org/wiki/Color_difference#CIEDE2000

func generateBackgroundVariations(backGround colorful.Color, numVariations int, ΔETarget float64, LStep float64, direction adjustmentDirection) (variations []colorful.Color, err error) {
	variations = make([]colorful.Color, numVariations)
	lastVariation := backGround
	for i := 0; i < numVariations; i++ {
		variation := lastVariation
		if i == 0 {
			variation = backGround
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
					return nil, fmt.Errorf("overflow L for variant %d", i)
				case l <= 0:
					return nil, fmt.Errorf("underflow L for variant %d", i)
				}
				variation = colorful.Lab(l, a, b)
				distance = lastVariation.DistanceCIEDE2000(variation)
			}
		}
		if !variation.IsValid() {
			variation = variation.Clamped()
		}
		variations[i] = variation
		lastVariation = variation
	}
	return variations, nil
}
