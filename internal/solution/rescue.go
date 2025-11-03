package solution

func Rescue(roofSize int, positions []int, maxCover int) int {
	chickenCount := 0

	if len(positions) == 1 {
		if roofSize != 0 && maxCover == 0 {
			return maxCover + 1
		}
		return maxCover
	}

	coverRange := roofSize + positions[0]
	if positions[1] > coverRange {
		positions = positions[1:]
	} else {
		for i := range roofSize {
			if i >= len(positions) {
				break
			}

			if positions[i] < coverRange {
				chickenCount++
			}
		}
		positions = positions[1:]
	}

	if maxCover < chickenCount {
		maxCover = chickenCount
	}

	return Rescue(roofSize, positions, maxCover)
}
