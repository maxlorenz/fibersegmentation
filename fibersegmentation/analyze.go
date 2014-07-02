package fibersegmentation

import (
"math"
)

func FiberLength(fiber []Pixel) float64 {
	lastPixel := fiber[0]
	length := 0.0

	for _, pixel := range fiber {
		if pixel.X != lastPixel.X {
			length = length + math.Sqrt(1.0 + float64(lastPixel.Y - pixel.Y) * float64(lastPixel.Y - pixel.Y))
			lastPixel = pixel
		}
	}

	return length
}
