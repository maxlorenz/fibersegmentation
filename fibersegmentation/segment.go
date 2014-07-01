package fibersegmentation

import (
	"image"
	"image/color"
	"image/draw"
)

func Segment(src image.Image, high float32, low uint8) image.Image {

	// create new image with same size
	bounds := src.Bounds()
	m := image.NewRGBA(bounds)

	// copy content to new image
	draw.Draw(m, bounds, src, bounds.Min, draw.Src)

	height := bounds.Max.Y
	width := bounds.Max.X

	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := src.At(x, y).RGBA()
			if (uint8(r) <= low) || (float32(r)/float32(g) <= high) {
				m.Set(x, y, color.NRGBA{0, 0, 0, 255})
			}
		}
	}

	return m
}
