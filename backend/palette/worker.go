package palette

import (
	"fmt"
	"github.com/mdesenfants/gokmeans"
	"github.com/nfnt/resize"
	"github.com/sepal/image_palette/backend/models"
	"image"
	"log"
)

// A worker calculates the 5 most dominant colors in an image based on the k-means algorithm.
// github.com/mdesenfants/gokmeans is used to caculate the kmeans.

// MAX_HEIGHT defines maximal height to which the an image will be resized for the k-means calculation.
const MAX_HEIGHT = 200

// Palette is a slice with the dominant colors for an image.
type Palette [5][3]uint8

// ToHex returns the palette as a hex color code (e.g. #112233)
func (p Palette) ToHex() (res models.Palette) {
	for i, val := range p {
		r := val[0]
		g := val[1]
		b := val[2]
		res[i] = fmt.Sprintf("#%X%X%X", r, g ,b)
	}
	return res
}

// calcPalette returns the dominant colors the given image.
func calcPalette(img image.Image) (Palette, error) {
	var palette Palette

	w := img.Bounds().Max.X - img.Bounds().Min.X
	h := img.Bounds().Max.Y - img.Bounds().Min.Y

	// Create a slice of observations, with the size of the image
	observations := make([]gokmeans.Node, w*h)

	// The k-means lib needs a continuous slice of observations(in our case pixels) instead of two dimensional
	// slice. This loop creates this slice, idx keeps the index for the observations.
	idx := 0
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			// Get the pixel at the current location, ignore any alpha value.
			r, g, b, _ := img.At(x, y).RGBA()

			// We only work with 8 bit per color for now.
			rF := float64(r) / 255.0
			gF := float64(g) / 255.0
			bF := float64(b) / 255.0

			observations[idx] = gokmeans.Node{rF, gF, bF}
			idx++
		}
	}

	// Calculate 5 centroids of the observations, which are in our case are the dominant colors in max. 10 rounds.
	success, centroids := gokmeans.Train(observations, 5, 10)

	if !success {
		return palette, fmt.Errorf("Could not calculate centroids.")
	}

	for idx, node := range centroids {
		palette[idx][0] = uint8(node[0])
		palette[idx][1] = uint8(node[1])
		palette[idx][2] = uint8(node[2])
	}
	return palette, nil
}

// Worker routine to calculate and save the palette of an image.
func Worker(image models.Image) {
	log.Printf("Starting calculating color scheme for file %v", image.Filename)
	img, err := image.GetImage()

	if err != nil {
		log.Println("Could not open image %v, due to: %v.", image.Filename, err)
	}

	// Resize the image, to accelerate the calculation.
	// todo: test different interpolation functions
	thumbnail := resize.Resize(0, MAX_HEIGHT, *img, resize.NearestNeighbor)

	palette, err := calcPalette(thumbnail)

	if err != nil {
		log.Fatalf("Could calculate palette for %v, because of: %v", image.Filename, err)
	}

	err = image.SavePalette(palette.ToHex())

	if err != nil {
		log.Fatalf("Error while trying to save palette to database.")
	}

	log.Printf("Finished calculating color scheme for file %v")
}
