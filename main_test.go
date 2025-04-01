package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistances(t *testing.T) {
	// given
	col1 := []int{3, 4, 2, 1, 3, 3}
	col2 := []int{4, 3, 5, 3, 9, 3}

	// when
	distances := calculateDistances(col1, col2)

	// then
	assert.Equal(t, 11, distances)
}
