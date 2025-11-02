package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		positions := []int{2, 5, 10, 12, 15}
		roofSize := 5

		maxCover := rescueRecursively(roofSize, positions)

		assert.Equal(t, 2, maxCover)
	})
}
