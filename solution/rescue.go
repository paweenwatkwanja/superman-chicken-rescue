package solution

func Rescue(roofSize int, positions []int, maxCoveredChicken int) int {
	chickenCount := 0

	if len(positions) == 1 {
		if maxCoveredChicken == 0 {
			return 1
		}
		return maxCoveredChicken
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

	if maxCoveredChicken < chickenCount {
		maxCoveredChicken = chickenCount
	}

	return Rescue(roofSize, positions, maxCoveredChicken)
}
