package main

import "log"

/* define a circle */
type GroupData struct {
	FirstValue  int
	SecondValue int
	ShapeType   string
}

/* define a method for circle */
func (groupData GroupData) area() int {
	switch groupData.ShapeType {
	case "persegi":
		return groupData.FirstValue * groupData.SecondValue
	case "segitiga":
		return (groupData.FirstValue * groupData.SecondValue) / 2
	default:
		return 0
	}
}

func main() {
	groupData := GroupData{
		FirstValue:  5,
		SecondValue: 10,
		ShapeType:   "persegi",
	}
	log.Printf("groupData area: %v", groupData.area())
}
