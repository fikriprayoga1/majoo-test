package main

import "log"

func main() {
	sourceValue := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	result := BubbleSort(sourceValue, "asc")
	log.Printf("logInfo : Result => %v", result)
}

func BubbleSort(array []float64, sortType string) []float64 {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if sortType == "asc" {
				if array[j] > array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			} else {
				if array[j] < array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}

		}
	}
	return array
}
