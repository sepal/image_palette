package color_converter

import (
	"testing"
)

func TestRGB_ToHSL(t *testing.T) {
	// Convert red to hsl
	expected := HSL{0, 1, 0.5}
	rgb := RGB{255, 0, 0}
	hsl := rgb.ToHSL()

	if hsl != expected {
		t.Errorf("Wrong convertion from rgb %v to hsl %v, expected %v", rgb, hsl, expected)
	}

	// Convert cyan to hsl
	expected = HSL{0.5, 1, 0.5}
	rgb = RGB{0, 255, 255}
	hsl = rgb.ToHSL()

	if hsl != expected {
		t.Errorf("Wrong convertion from rgb %v to hsl %v, expected %v", rgb, hsl, expected)
	}

	// Convert white to hsl
	expected = HSL{0, 0, 1}
	rgb = RGB{255, 255, 255}
	hsl = rgb.ToHSL()

	if hsl != expected {
		t.Errorf("Wrong convertion from rgb %v to hsl %v, expected %v", rgb, hsl, expected)
	}

	// Convert black to hsl
	expected = HSL{0, 0, 0}
	rgb = RGB{0, 0, 0}
	hsl = rgb.ToHSL()

	if hsl != expected {
		t.Errorf("Wrong convertion from rgb %v to hsl %v, expected %v", rgb, hsl, expected)
	}
}
