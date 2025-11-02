package main

import "fmt"

func rescueRecursively(roofSize int, positions []int) int {
	chickenCount := 0
	maxCover := 0

	if len(positions) == 1 {
		return 0
	}

	coverRange := roofSize + positions[0]
	if positions[1] >= coverRange {
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

	fmt.Printf("positions : %v; max cover : %v\n", positions, maxCover)
	return rescueRecursively(roofSize, positions) + maxCover
}
