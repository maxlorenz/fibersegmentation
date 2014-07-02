package fibersegmentation

func (p1 Pixel) LessThan(p2 Pixel) bool {
	if p1.X < p2.X {
		return true
	}

	if p1.X == p2.X && p1.Y < p2.Y {
		return true
	}

	return false
}

func SortFibers(unsortedFibers [][]Pixel) [][]Pixel {
	sortedFibers := unsortedFibers

	for i, fiber := range unsortedFibers {
		sortedFibers[i] = sortPixels(fiber)
	}

	return sortedFibers

}

func sortPixels(fiber []Pixel) []Pixel {
	if len(fiber) < 2 {
		return fiber
	}

	left, right := 0, len(fiber)-1

	fiber[0], fiber[right] = fiber[right], fiber[0]

	for i := range fiber {
		if fiber[i].LessThan(fiber[right]) {
			fiber[i], fiber[left] = fiber[left], fiber[i]
			left++
		}
	}

	fiber[left], fiber[right] = fiber[right], fiber[left]

	sortPixels(fiber[:left])
	sortPixels(fiber[left+1:])

	return fiber
}
