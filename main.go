package main

import "fmt"

func main() {
	roofSize := 5
	positions := []int{2, 5, 10, 12, 15}
	// roofSize := 10
	// positions := []int{1, 11, 30, 34, 35, 37}

	var maxCover int
	fmt.Println(rescueRecursively(roofSize, positions, maxCover))
}
