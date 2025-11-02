package main

import "fmt"

func main() {
	// roofSize := 5
	// positions := []int{2, 5, 10, 12, 15}
	roofSize := 10
	positions := []int{1, 11, 30, 34, 35, 37}
	n := len(positions)

	chickenInCover := 0
	maxChicken := 0
	for i := range n {
		start := positions[i]
		cover := start + roofSize

		for _, pst := range positions {
			if pst < cover && pst >= start {
				chickenInCover++
			}
		}
		if maxChicken < chickenInCover {
			maxChicken = chickenInCover
		}
		chickenInCover = 0
	}
	fmt.Println(maxChicken)
}
