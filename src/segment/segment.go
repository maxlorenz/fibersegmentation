package segment

import(
	"image"
	"image/color"
	"image/draw"
)

type Fiber struct {
	Original, Segmented image.Image
}

func (self *Fiber) Segment(src image.Image, high float64, low uint8) image.Image {

	self.Original = src

	// create new image with same size
	bounds := self.Original.Bounds()
	m := image.NewRGBA(bounds)

	// copy content to new image
	draw.Draw(m, bounds, self.Original, bounds.Min, draw.Src)

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

	self.Segmented = m

	return self.Segmented
}