package fibersegmentation

import(
	"image"
)

type Pixel struct {
	X, Y int
}

func ConnectedComponents(src image.Image, high float64, low uint8) map[Pixel]int {

	height := src.Bounds().Max.Y
	width := src.Bounds().Max.X

	regionCount := 0

	pixelMarks := map[Pixel]int {}
	equivalentRegions := map[int]int {}

	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := src.At(x, y).RGBA()
			if (uint8(r) <= low) || (float64(r)/float64(g) <= high) {

				// Check north and west pixel for 4-connectivity
				w, wExists := pixelMarks[Pixel{x - 1, y}]
				n, nExists := pixelMarks[Pixel{x, y - 1}]


				// First check if both pixels are not marked and give current pixel a new one
				if !wExists && !nExists {
					pixelMarks[Pixel{x, y}] = regionCount
					regionCount++
					equivalentRegions[regionCount] = regionCount
				}

				// One pixel -> current gets same marker
				if wExists && !nExists {
					pixelMarks[Pixel{x, y}] = w
				}

				if !wExists && nExists {
					pixelMarks[Pixel{x, y}] = n
				}

				// Are both pixels marked?
				if wExists && nExists {
					if n != w { // Different marker -> make an entry in the equivalence map
						parent := n
						for equivalentRegions[parent] != parent {
							parent = equivalentRegions[parent] // Simple union-find
						}

						equivalentRegions[w] = parent
					}

					pixelMarks[Pixel{x, y}] = n
				}

			}
		}
	}

	for pixel := range pixelMarks {
		pixelMarks[pixel] = equivalentRegions[pixelMarks[pixel]]
	}

	return pixelMarks
}