package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea(t *testing.T) {
	groupData := GroupData{
		FirstValue:  5,
		SecondValue: 10,
		ShapeType:   "persegi",
	}
	result := groupData.area()
	expected := 50
	assert.Equal(t, expected, result)
}
