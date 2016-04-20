package color_converter

import (
	"fmt"
	"math"
)

type HSV struct {
	H, S, V float64
}

func (rgb *RGB) ToHSV() HSV {
	r, g, b := rgb.Normalized()

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	v := max

	d := max - min

	var s float64
	if max == 0 {
		s = 0
	} else {
		s = d / max
	}

	if max == min {
		return HSV{0, s, v}
	}

	var h float64
	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}

	h /= 6

	return HSV{h, s, v}
}

func (hsv *HSV) String() string {
	return fmt.Sprintf("%v°  %v%%  %v%%", uint8(hsv.H*360), uint8(hsv.S*100), uint8(hsv.V * 100))
}
