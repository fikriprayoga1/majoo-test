package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescendingSort(t *testing.T) {
	sourceValue := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	result := BubbleSort(sourceValue, "dsc")
	expected := []float64{10, 9, 4, 3.3, 3, 0.2, 0, -5, -7}
	assert.Equal(t, expected, result)
}

func TestAscendingSort(t *testing.T) {
	sourceValue := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	result := BubbleSort(sourceValue, "asc")
	expected := []float64{-7, -5, 0, 0.2, 3, 3.3, 4, 9, 10}
	assert.Equal(t, expected, result)
}
