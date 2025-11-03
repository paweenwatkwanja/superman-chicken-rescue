package main

import "fmt"

func main() {
	var roofSize int
	var positions = []int{1}

	roofSize = 2
	// check count of position and n of input
	var maxCover int
	fmt.Println(rescueRecursively(roofSize, positions, maxCover))
}
