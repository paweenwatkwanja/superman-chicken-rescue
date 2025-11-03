package main

import "fmt"

func main() {
	var roofSize int
	var positions = []int{1}

	roofSize = 2
	// check count of position and n of input
	// if positions length = 0 no process
	// first position must not be 0
	var maxCover int
	fmt.Println(rescueRecursively(roofSize, positions, maxCover))
}
