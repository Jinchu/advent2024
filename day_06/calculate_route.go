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

func getGridSize(inputLines []string) input.Coordinates {
	var size input.Coordinates
	size.Y = len(inputLines)
	size.X = len(inputLines[0])
	return size
}

func sameCoordinates(comp1 input.Coordinates, comp2 input.Coordinates) bool {
	if comp1.X == comp2.X && comp1.Y == comp2.Y {
		return true
	} else {
		return false
	}
}

func travelNorth(
	mapSize input.Coordinates, labMap []input.Coordinates,
	position input.Coordinates, direction Direction) input.Coordinates {
	var trail []input.Coordinates

	previousPosition := position
	for y := position.Y; y < mapSize.Y; {

		position.Y = y
		for _, block := range labMap {
			if sameCoordinates(block, position) {
				// Hitting a block
				fmt.Printf("Found a block at position %v\n", previousPosition)
				return previousPosition
			}
		}
		trail = append(trail, position) // TODO: this must be set or something like that
		previousPosition = position
		if direction == north {
			y--
		}
	}
	// TODO: found exit
	fmt.Printf("The guard will exit here %v\n", previousPosition)

	return previousPosition
}

func travel(
	mapSize input.Coordinates, labMap []input.Coordinates,
	position input.Coordinates, direction Direction) {
	debug := false
	switch direction {
	case north:
		travelNorth(mapSize, labMap, position, direction)
	case east:
	case south:
	case west:
	}

	if debug {
		fmt.Println(labMap)
	}
}

func CalculateRoute() {
	total := 0

	inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	// inputLines := input.GetInputV2("../downloads/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")
	mapSize := getGridSize(inputLines)

	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v guards in total.\n", len(startingPoint))
	fmt.Printf("The total size of the map is %v by %v\n", mapSize.X, mapSize.Y)

	travel(mapSize, blockCoordinates, startingPoint[0], north)
	fmt.Printf("The lab guard will visit %v distinct positions\n", total)
}
