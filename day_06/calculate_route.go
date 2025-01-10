package day06

import (
	"advent/internal/input"
	"fmt"
)

func CalculateRoute() {
	total := 0

	inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	// inputLines := input.GetInputV2("./day_06/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")
	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v blocks in total.\n", len(startingPoint))

	fmt.Printf("The lab guard will visit %v distinct positions\n", total)
}
