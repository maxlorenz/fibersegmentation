package fibersegmentation

import(
	"image"
)

func ConnectedComponents(src image.Image, high float64, low uint8) map[Pixel]int {

	height := src.Bounds().Max.Y
	width := src.Bounds().Max.X

	regionCount := 0

	markedMap := map[Pixel]int {}
	equivalentRegions := map[int]int {}

	// test each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := src.At(x, y).RGBA()
			if (uint8(r) <= low) || (float64(r)/float64(g) <= high) {

				// Check north and west pixel for 4-connectivity
				w, wexists := markedMap[Pixel{x - 1, y}]
				n, nexists := markedMap[Pixel{x, y - 1}]


				// First check if both pixels are not marked and give current pixel a new one
				if !wexists && !nexists {
					markedMap[Pixel{x, y}] = regionCount
					regionCount++
					equivalentRegions[regionCount] = regionCount
				}

				// One pixel -> current gets same marker
				if wexists && !nexists {
					markedMap[Pixel{x, y}] = w
				}

				if !wexists && nexists {
					markedMap[Pixel{x, y}] = n
				}

				// Are both pixels marked?
				if wexists && nexists {
					if n != w { // Different marker -> make an entry in the equivalence map
						temp := n
						for equivalentRegions[temp] != temp {
							temp = equivalentRegions[temp]
						}

						equivalentRegions[w] = temp
					}

					markedMap[Pixel{x, y}] = n
				}

			}
		}
	}

	for pixel := range markedMap {
		markedMap[pixel] = equivalentRegions[markedMap[pixel]]
	}

	return markedMap
}