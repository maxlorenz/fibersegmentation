package segment

import(
	"image"
	"image/color"
	"image/draw"
)

type SegmentImage struct {
	Original, Segmented image.Image
}

func (self *SegmentImage) Segment(high float64, low uint8) {

	// create new image with same size
	bounds := self.Original.Bounds()
	m := image.NewRGBA(bounds)

	// copy content to new image
	draw.Draw(m, bounds, self.Original, bounds.Min, draw.Src)

	// test each pixel
	for y := 0; y < m.Bounds().Max.Y; y++ {
		for x := 0; x < m.Bounds().Max.X; x++ {
			r, g, _, _ := m.At(x, y).RGBA()
			if !(uint8(r) > low) || !(float64(r)/float64(g) > high) {
					m.Set(x, y, color.NRGBA {0, 0, 0, 255})
			}
		}
	}

	self.Segmented = m
}