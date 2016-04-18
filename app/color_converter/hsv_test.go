package color_converter

import "testing"

func TestRGB_ToHSV(t *testing.T) {
	// Convert red to hsv
	expected := HSV{0, 1, 1}
	rgb := RGB{255, 0, 0}
	hsv := rgb.ToHSV()

	if hsv != expected {
		t.Errorf("Wrong convertion from rgb %v to hsv %v, expected %v", rgb, hsv, expected)
	}

	// Convert cyan to hsv
	expected = HSV{0.5, 1, 1}
	rgb = RGB{0, 255, 255}
	hsv = rgb.ToHSV()

	if hsv != expected {
		t.Errorf("Wrong convertion from rgb %v to hsv %v, expected %v", rgb, hsv, expected)
	}

	// Convert white to hsv
	expected = HSV{0, 0, 1}
	rgb = RGB{255, 255, 255}
	hsv = rgb.ToHSV()

	if hsv != expected {
		t.Errorf("Wrong convertion from rgb %v to hsv %v, expected %v", rgb, hsv, expected)
	}

	// Convert black to hsv
	expected = HSV{0, 0, 0}
	rgb = RGB{0, 0, 0}
	hsv = rgb.ToHSV()

	if hsv != expected {
		t.Errorf("Wrong convertion from rgb %v to hsv %v, expected %v", rgb, hsv, expected)
	}
}
