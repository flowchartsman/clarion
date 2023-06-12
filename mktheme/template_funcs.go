package main

import (
	"fmt"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// these functions are used in the templates themselves to output colors in the
// appropriate format for the theme. It's important to remember that all of the
// color functions need to be mapped to their appropriate formats using pipes.
// For example, if you want the hex color for bg 0, it would be {{bg 0|hex}},
// whereas for the 256 color terminal version, you would want {{bg 0|term}}

var errColor = colorful.Hsl(300, 1, 0.5) // ugly ol' magenta

func tmplBG(table colorTable) func(int) colorful.Color {
	return func(offset int) colorful.Color {
		if offset < 0 || offset >= len(table.bg) {
			themeLogErr("bg offset out of range: %d", offset)
			return errColor
		}
		return table.bg[offset]
	}
}

func tmplColor(colorName string, table colorTable) func(string) colorful.Color {
	return func(contrastLevel string) colorful.Color {
		switch contrastLevel {
		case "min":
			return table.conceptColors[colorName].min
		case "enhanced":
			return table.conceptColors[colorName].enhanced
		case "ui":
			return table.conceptColors[colorName].ui
		case "base":
			return table.conceptColors[colorName].base
		default:
			themeLogErr(`invalid contrast level %q for color %q, must be one of [min,enhanced,ui]`, contrastLevel, colorName)
			return errColor
		}
	}
}

func tmplAnsiColor(colorName string, table colorTable) func() colorful.Color {
	return func() colorful.Color {
		return table.ansiColors[colorName]
	}
}

func tmplFG(table colorTable) func() colorful.Color {
	return func() colorful.Color {
		return table.fg
	}
}

func tmplLightFG(table colorTable) func() colorful.Color {
	return func() colorful.Color {
		return table.lightFg
	}
}

func tmplUIFG(table colorTable) func() colorful.Color {
	return func() colorful.Color {
		return table.ui
	}
}

func tmpl2Hex(color colorful.Color) string {
	if !color.IsValid() {
		return color.Clamped().Hex()
	}
	return color.Hex()
}

func tmpl2HexTrunc(color colorful.Color) string {
	return tmpl2Hex(color)[1:]
}

func tmpl2HexAlpha(transparencyPct float64, color colorful.Color) string {
	hexStr := tmpl2Hex(color)
	return fmt.Sprintf("%s%02x", hexStr, int(math.Round(float64(transparencyPct)*255.0/100.0)))
}

func tmpl2Term(color colorful.Color) string {
	if !color.IsValid() {
		return toTerminal(color.Clamped())
	}
	return toTerminal(color)
}
