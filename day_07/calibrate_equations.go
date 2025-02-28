package day07

import (
	"advent/internal/input"
	"fmt"
)

func GetTotalCalibration() {
	inputLines := input.GetInputV2("./day_07/test-input-1.txt")

	for _, line := range inputLines {
		fmt.Println(line)
	}
}
