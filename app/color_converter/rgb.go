package color_converter

import "fmt"

// RGB color space with values from 0 - 255.
type RGB struct {
	R, G, B uint8
}

func (rgb *RGB) Normalized() (r, g, b float64) {
	r = float64(rgb.R) / 255
	g = float64(rgb.G) / 255
	b = float64(rgb.B) / 255
	return r, g, b
}

func (rgb *RGB) String() string {
	return fmt.Sprintf("0x%x%x%x", rgb.R, rgb.G, rgb.B)
}