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

func travelNorthSouth(
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

		switch direction {
		case north:
			y--
		case south:
			y++
		default:
			panic("unexpected direction")
		}

	}
	// TODO: found exit
	fmt.Printf("The guard will exit here %v\n", previousPosition)

	return previousPosition
}

func travelEastWest(
	mapSize input.Coordinates, labMap []input.Coordinates,
	position input.Coordinates, direction Direction) input.Coordinates {
	var trail []input.Coordinates

	previousPosition := position
	for x := position.X; x < mapSize.X; {
		position.X = x

		for _, block := range labMap {
			if sameCoordinates(block, position) {
				fmt.Printf("Found a block at position %v\n", previousPosition)
				return previousPosition
			}
		}
		trail = append(trail, position) // TODO: this must be a set or something
		previousPosition = position

		switch direction {
		case east:
			x++
		case west:
			x--
		default:
			panic("unexpected direction")
		}
	}

	return previousPosition
}

func travel(
	mapSize input.Coordinates, labMap []input.Coordinates,
	position input.Coordinates, direction Direction) input.Coordinates {
	debug := false
	var updatedPosition input.Coordinates
	switch direction {
	case north:
		updatedPosition = travelNorthSouth(mapSize, labMap, position, direction)
	case east:
		updatedPosition = travelEastWest(mapSize, labMap, position, direction)
	case south:
		updatedPosition = travelNorthSouth(mapSize, labMap, position, direction)
	case west:
		updatedPosition = travelEastWest(mapSize, labMap, position, direction)
	default:
		panic("unexpected direction")
	}

	if debug {
		fmt.Println(labMap)
	}

	return updatedPosition
}

func guardNavigation(
	mapSize input.Coordinates, blockMap []input.Coordinates,
	position input.Coordinates) {
	i := 0
	j := 0
	allDirections := [...]Direction{north, east, south, west}

	for true {
		if j > 512 {
			panic("oops")
		}

		position = travel(mapSize, blockMap, position, allDirections[i])
		i++
		j++

		if i >= len(allDirections) {
			i = 0
		}

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

	guardNavigation(mapSize, blockCoordinates, startingPoint[0])
	fmt.Println("-----")

	position := travel(mapSize, blockCoordinates, startingPoint[0], north)
	position = travel(mapSize, blockCoordinates, position, east)
	position = travel(mapSize, blockCoordinates, position, south)
	position = travel(mapSize, blockCoordinates, position, west)
	fmt.Printf("The lab guard will visit %v distinct positions\n", total)
}
