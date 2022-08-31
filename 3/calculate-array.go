package main

import "log"

func main() {
	result := CalculateArray(5, 8, 7)
	log.Printf("logInfo : Result => %v", result)
}

func CalculateArray(firstValue int, secondValue int, xValue int) []int {
	var result []int

	differenceValue := secondValue - firstValue

	for i := 0; i < xValue; i++ {
		result = append(result, firstValue)
		firstValue += differenceValue
	}

	return result
}
