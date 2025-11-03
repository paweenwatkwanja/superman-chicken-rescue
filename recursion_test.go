package main

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
			actualMaxCover := rescueRecursively(tc.roofSize, tc.positions, 0)

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
		{"each position is too far from each other", 5, []int{1, 6, 11, 16, 22}, 1},

		{"no space between positions and roof size 1", 1, []int{1, 2, 3, 4, 5}, 1},
		{"no space between positions with roof size 2", 2, []int{1, 2, 3, 4, 5}, 2},

		{"no space between positions and roof size 0", 0, []int{1, 2, 3, 4, 5}, 0},

		{"only one position and roof size 5", 5, []int{1}, 1},
	}

	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %v: %v", i+1, tc.testCaseName)
		t.Run(testName, func(t *testing.T) {
			actualMaxCover := rescueRecursively(tc.roofSize, tc.positions, 0)

			assert.Equal(t, tc.expectedMaxCover, actualMaxCover)
		})
	}
}
