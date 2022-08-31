package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	result := CalculateArray(5, 8, 7)
	expected := []int{5, 8, 11, 14, 17, 20, 23}
	assert.Equal(t, expected, result)
}

func TestCalculate2(t *testing.T) {
	result := CalculateArray(2, 4, 5)
	expected := []int{2, 4, 6, 8, 10}
	assert.Equal(t, expected, result)
}
