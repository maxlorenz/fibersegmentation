package segment

import(
	"image"
	"fmt"
)

type Pixel struct {
	X, Y int
}

type Fiber struct {
	Pixels []Pixel
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

// Needs to be rewritten as Pixels come unordered and form new Fibers too often!
// Arrays are segmented because of that mechanism
func (self *ImageAnalytics) SeparateFibers() []Fiber {
	result := []Fiber {Fiber {}}

	for i := range self.Fibers {
		
		currentPixel := self.Fibers[i]
		pixelBelongsToFiber := false

		for j := range result {
			if currentPixel.BelongsTo(result[j]) {
				result[j].Pixels = append(result[j].Pixels, currentPixel)
				pixelBelongsToFiber = true
			}
		}

		if !pixelBelongsToFiber {
			fmt.Println("%v", result)
			result = append(result, Fiber {[]Pixel {currentPixel}})
		}
	}

	return result
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