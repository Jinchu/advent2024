package day06

import (
	"advent/internal/input"
	"fmt"
)

type Direction int64

const (
	north Direction = iota
	east
	south
	west
)

func getGridSize(inputLines []string) string {
	return "LOL"
}

func sameCoordinates(comp1 input.Coordinates, comp2 input.Coordinates) bool {
	if comp1.X == comp2.X && comp1.Y == comp2.Y {
		return true
	} else {
		return false
	}
}

func travelNorth(labMap []input.Coordinates, position input.Coordinates) input.Coordinates {
	var trail []input.Coordinates

	previousPosition := position
	for y := position.Y; y < 7; y++ {
		position.Y = y
		for _, block := range labMap {
			if sameCoordinates(block, position) {
				// Hitting a block
				return previousPosition
			}
		}
		trail = append(trail, position) // TODO: this must be set or something like that
		previousPosition = position
	}

	return previousPosition
}

func travel(labMap []input.Coordinates, position input.Coordinates, direction Direction) {

	switch direction {
	case north:
	case east:
	case south:
	case west:
	}
	fmt.Println(labMap)
}

func CalculateRoute() {
	total := 0

	inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	// inputLines := input.GetInputV2("./day_06/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")
	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v blocks in total.\n", len(startingPoint))

	travel(blockCoordinates, startingPoint[0], north)
	fmt.Printf("The lab guard will visit %v distinct positions\n", total)
}
