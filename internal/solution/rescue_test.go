package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursionWithProvidedTestCases(t *testing.T) {
	testCases := []struct {
		roofSize         int
		positions        []int
		expectedMaxCover int
	}{
		{5, []int{2, 5, 10, 12, 15}, 2},
		{10, []int{1, 11, 30, 34, 35, 37}, 4},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v", i+1)
		t.Run(testName, func(t *testing.T) {
			actualMaxCover := Rescue(tc.roofSize, tc.positions, 0)

			assert.Equal(t, tc.expectedMaxCover, actualMaxCover)
		})
	}
}

func TestRecursionWithAdditionalTestCases(t *testing.T) {
	testCases := []struct {
		testCaseName     string
		roofSize         int
		positions        []int
		expectedMaxCover int
	}{
		{"many positions but the next position of each is beyond the roof", 5, []int{1, 6, 11, 16, 22}, 1},

		{"many positions next to each other with roof size 1", 1, []int{1, 2, 3, 4, 5}, 1},
		{"many positions next to each other with roof size 2", 2, []int{1, 2, 3, 4, 5}, 2},

		{"many positions next to each other with roof size cover the last position", 5, []int{1, 2, 3, 4, 5}, 5},
		{"many positions next to each other with roof size cover beyond last position", 6, []int{1, 2, 3, 4, 5}, 5},

		{"one position with roof size greater than 0", 5, []int{1}, 1},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v: %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			actualMaxCover := Rescue(tc.roofSize, tc.positions, 0)

			assert.Equal(t, tc.expectedMaxCover, actualMaxCover)
		})
	}
}
