package main

import (
	"fmt"

	"github.com/paweenwatkwanja/superman-chicken-rescue/solution"
	"github.com/paweenwatkwanja/superman-chicken-rescue/utils"
)

func main() {
	rescueInput := utils.GetInputFromTerminal()
	err := utils.ValidateInput(rescueInput)
	if err != nil {
		panic(err)
	}

	var maxCoveredChicken int
	maxCoveredChicken = solution.Rescue(rescueInput.RoofSize, rescueInput.Positions, maxCoveredChicken)
	fmt.Println(maxCoveredChicken)
}
