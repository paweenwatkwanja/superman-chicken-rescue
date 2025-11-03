package main

import (
	"fmt"

	"github.com/paweenwatkwanja/superman-chicken-rescue/internal/solution"
	"github.com/paweenwatkwanja/superman-chicken-rescue/utils"
)

func main() {
	rescueInput := utils.GetInputFromTerminal()
	err := utils.ValidateInput(rescueInput)
	if err != nil {
		panic(err)
	}

	var maxCover int
	maxCover = solution.Rescue(rescueInput.RoofSize, rescueInput.Positions, maxCover)
	fmt.Println(maxCover)
}
