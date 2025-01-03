package day_05

import (
	"advent/internal/input"
	"fmt"
)

func PrintOrder() {
	inputLines := input.GetInputV2("./day_05/test-input-1.txt")
	//	inputLines := input.GetInputV2("./day_04/input-day6.txt")
	for _, line := range inputLines {
		fmt.Println(line)
	}
}
