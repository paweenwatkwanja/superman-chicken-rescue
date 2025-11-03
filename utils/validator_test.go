package utils

import (
	"errors"
	"fmt"
	"testing"

	"github.com/paweenwatkwanja/superman-chicken-rescue/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateInput(t *testing.T) {
	testCases := []struct {
		rescueInput *models.RescueInput
		err         error
	}{
		{&models.RescueInput{
			NumberOfChickens: 1,
			RoofSize:         1,
			Positions:        []int{1},
		}, nil},
		{&models.RescueInput{
			NumberOfChickens: 0,
			RoofSize:         1,
			Positions:        []int{1},
		}, errors.New("error : NumberOfChickens input out of range")},
		{&models.RescueInput{
			NumberOfChickens: 1,
			RoofSize:         10000000,
			Positions:        []int{1},
		}, errors.New("error : RoofSize input out of range")},
		{&models.RescueInput{
			NumberOfChickens: 1,
			RoofSize:         1,
			Positions:        []int{},
		}, errors.New("error : number of chickens must be equal to the length of positions")},
		{&models.RescueInput{
			NumberOfChickens: 1,
			RoofSize:         1,
			Positions:        []int{0},
		}, errors.New("error : the position must not start with 0")},
		{&models.RescueInput{
			NumberOfChickens: 2,
			RoofSize:         1,
			Positions:        []int{1, 1000000001},
		}, errors.New("error : the position must not be over 1,000,000,000")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			err := ValidateInput(tc.rescueInput)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}

func TestValidateLength(t *testing.T) {
	testCases := []struct {
		length int
		err    error
	}{
		{0, errors.New("input out of range")},
		{1, nil},
		{1000000, nil},
		{1000001, errors.New("input out of range")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			err := validateLength(tc.length)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}

func TestValidatePositionLength(t *testing.T) {
	testCases := []struct {
		positionLength   int
		numberOfChickens int
		err              error
	}{
		{1, 1, nil},
		{1, 2, errors.New("number of chickens must be equal to the length of positions")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			err := validatePositionLength(tc.positionLength, tc.numberOfChickens)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}

func TestValidateFirstPosition(t *testing.T) {
	testCases := []struct {
		firstPosition int
		err           error
	}{
		{1, nil},
		{0, errors.New("the position must not start with 0")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			err := validateFirstPosition(tc.firstPosition)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}

func TestValidateLastPosition(t *testing.T) {
	testCases := []struct {
		lastPosition int
		err          error
	}{
		{1000000000, nil},
		{1000000001, errors.New("the position must not be over 1,000,000,000")},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			err := validateLastPosition(tc.lastPosition)

			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.err, nil)
			}
		})
	}
}
