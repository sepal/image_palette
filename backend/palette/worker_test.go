package palette

import (
	"testing"
)

func TestPalette_ToHex(t *testing.T) {
	type testColor struct {
		hex string
		rgb [3]uint8
	}

	colors := make([]testColor, 5)

	colors[0] = testColor{"#FF0000", [3]uint8{255, 0, 0}}
	colors[1] = testColor{"#00FF00", [3]uint8{0, 255, 0}}
	colors[2] = testColor{"#0000FF", [3]uint8{0, 0, 255}}
	colors[3] = testColor{"#FFFF00", [3]uint8{255, 255, 0}}
	colors[4] = testColor{"#FFFFFF", [3]uint8{255, 255, 255}}

	var p Palette

	for i, c := range colors {
		p[i] = c.rgb
	}

	hex := p.ToHex()

	for i, c := range colors {
		if hex[i] != c.hex {
			t.Errorf("ToHex color %v doesn't match %v.", hex[i], c.hex)
		}
	}
}
