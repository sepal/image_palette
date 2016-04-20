package color_converter

import (
	"math"
	"fmt"
)

// HSL color space with values from 0 - 1.
type HSL struct {
	H, S, L float64
}

func (rgb *RGB) ToHSL() HSL {
	r, g, b := rgb.Normalized()

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l := (max + min) / 2

	if max == min {
		return HSL{0, 0, l}
	}
	var h, s float64
	delta := max - min

	if l > 0.5 {
		s = delta / (2 - max - min)
	} else {
		s = delta / (max + min)
	}

	switch max {
	case r:
		h = (g - b) / delta
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/delta + 2
	case b:
		h = (r-g)/delta + 4
	}

	h /= 6

	return HSL{h, s, l}
}

func (hsl *HSL) String() string {
	return fmt.Sprintf("%v°  %v%%  %v%%", uint8(hsl.H*360), uint8(hsl.S*100), uint8(hsl.L * 100))
}
