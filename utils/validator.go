package utils

import (
	"errors"
	"fmt"

	"github.com/paweenwatkwanja/superman-chicken-rescue/models"
)

func ValidateInput(rescueInput *models.RescueInput) error {
	var err error

	err = validateLength(rescueInput.NumberOfChickens)
	if err != nil {
		return fmt.Errorf("error : NumberOfChickens %v", err.Error())
	}

	err = validateLength(rescueInput.RoofSize)
	if err != nil {
		return fmt.Errorf("error : RoofSize %v", err.Error())
	}

	positionLength := len(rescueInput.Positions)
	err = validatePositionLength(positionLength, rescueInput.NumberOfChickens)
	if err != nil {
		return fmt.Errorf("error : %v", err.Error())
	}

	err = validateFirstPosition(rescueInput.Positions[0])
	if err != nil {
		return fmt.Errorf("error : %v", err.Error())
	}

	lastPosition := rescueInput.Positions[rescueInput.NumberOfChickens-1]
	err = validateLastPosition(lastPosition)
	if err != nil {
		return fmt.Errorf("error : %v", err.Error())
	}
	return nil
}

func validateLength(length int) error {
	if length < 1 || length > 1000000 {
		return errors.New("input out of range")
	}

	return nil
}

func validatePositionLength(positionLength int, numberOfChickens int) error {
	if positionLength != numberOfChickens {
		return errors.New("number of chickens must be equal to the length of positions")
	}

	return nil
}

func validateFirstPosition(firstPosition int) error {
	if firstPosition < 1 {
		return errors.New("the position must not start with 0")
	}

	return nil
}

func validateLastPosition(lastPosition int) error {
	if lastPosition > 1000000000 {
		return errors.New("the position must not be over 1,000,000,000")
	}

	return nil
}
