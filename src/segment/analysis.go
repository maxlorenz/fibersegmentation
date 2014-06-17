package segment

import(
	"image"
	// "fmt"
	// "../graph"
)

type Pixel struct {
	X, Y int
}

type Fiber struct {
	Pixels []Pixel
}

type ImageAnalytics struct {
	Fibers []Pixel
	Table []Fiber
}



func ReadToMemory(src image.Image, high float64, low uint8) ImageAnalytics {

	analytics := new (ImageAnalytics)

	table := make([]Fiber, 0)

	height := src.Bounds().Max.Y
	width := src.Bounds().Max.X

	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := src.At(x, y).RGBA()
			if (uint8(r) > low) && (float64(r)/float64(g) > high) {

				pixel := Pixel {x, y}
				found := false

				for i := range table {
					if pixel.BelongsTo(table[i]) {
						table[i] = Fiber {append(table[i].Pixels[:], pixel)}
						found = true
					}
				}
				if found == false {
					table = append(table, Fiber {[]Pixel {pixel}})
				}

			}
		}
	}

	analytics.Table = table
	return *analytics
}

func (self *Pixel) BelongsTo(fiber Fiber) bool {

	nextPixels := []Pixel {
		Pixel {self.X + 1, self.Y},
		Pixel {self.X - 1, self.Y},
		Pixel {self.X, self.Y + 1},
		Pixel {self.X, self.Y - 1},
		Pixel {self.X - 1, self.Y - 1},
		Pixel {self.X + 1, self.Y - 1},
		Pixel {self.X - 1, self.Y + 1},
		Pixel {self.X + 1, self.Y + 1},
	}

	for i := range fiber.Pixels {
		currentPixel := fiber.Pixels[i]

		for j := range nextPixels {
			if currentPixel == nextPixels[j] {
				return true
			}
		}
	}

	return false
}