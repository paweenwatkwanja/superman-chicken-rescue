package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/paweenwatkwanja/superman-chicken-rescue/models"
)

func GetInputFromTerminal() *models.RescueInput {
	fmt.Println("Enter the input: ")
	var numberOfChickens, roofSize int
	_, err := fmt.Scan(&numberOfChickens, &roofSize)
	if err != nil {
		fmt.Println("error reading first line:", err)
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var positions []int
	inputs := strings.Fields(line)
	for _, input := range inputs {
		position, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("error converting '%v' to integer: %v\n", input, err)
			return nil
		}
		positions = append(positions, position)
	}

	return &models.RescueInput{
		NumberOfChickens: numberOfChickens,
		RoofSize:         roofSize, Positions: positions}
}
