package segment

import(
	"image"
)

type Pixel struct {
	X, Y int
}

type ImageAnalytics struct {
	Fibers []Pixel
}

func ReadToMemory(src image.Image, high float64, low uint8) ImageAnalytics {

	analytics := new (ImageAnalytics)

	height := src.Bounds().Max.Y
	width := src.Bounds().Max.X

	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := src.At(x, y).RGBA()
			if (uint8(r) > low) && (float64(r)/float64(g) > high) {
				analytics.Fibers = append(analytics.Fibers, Pixel {x, y})
			}
		}
	}

	return *analytics
}