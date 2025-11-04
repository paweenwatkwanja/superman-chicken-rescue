package main

import (
	"fmt"
	"time"

	"github.com/paweenwatkwanja/superman-chicken-rescue/solution"
	"github.com/paweenwatkwanja/superman-chicken-rescue/utils"
)

func main() {
	rescueInput := utils.GetInputFromTerminal()
	err := utils.ValidateInput(rescueInput)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	var maxCoveredChicken int
	maxCoveredChicken = solution.Rescue(rescueInput.RoofSize, rescueInput.Positions, maxCoveredChicken)
	fmt.Println(maxCoveredChicken)

	elapsed := time.Since(start)
	fmt.Printf("Duration: %v\n", elapsed)
}
