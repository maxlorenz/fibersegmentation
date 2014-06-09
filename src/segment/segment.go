package segment

import(
	"image"
	"image/color"
	"image/draw"
)

func Segment(src image.Image, high float64, low uint8) image.Image {

	// create new image with same size
	bounds := src.Bounds()
	m := image.NewRGBA(bounds)

	// copy content to new image
	draw.Draw(m, bounds, src, bounds.Min, draw.Src)

	height := m.Bounds().Max.Y
	width := m.Bounds().Max.X
	
	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := m.At(x, y).RGBA()
			if (uint8(r) <= low) || (float64(r)/float64(g) <= high) {
				m.Set(x, y, color.NRGBA {0, 0, 0, 255})
			}
		}
	}

	return m
}