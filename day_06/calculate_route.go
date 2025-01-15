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

// Calculate the travel in North-South direction. Returns true if the route exist the grid.
// Otherwise false
func (route *guardRoute) travelNorthSouth(labMap []input.Coordinates) (bool, *guardRoute) {
	var trail []input.Coordinates

	previousPosition := route.position
	for y := route.position.Y; y < route.mapSize.Y; {

		route.position.Y = y
		for _, block := range labMap {
			if input.SameCoordinates(block, route.position) {
				// Hitting a block
				fmt.Printf("Found a block at position %v\n", previousPosition)
				route.position = previousPosition
				return false, route
			}
		}
		trail = append(trail, route.position) // TODO: this must be set or something like that
		previousPosition = route.position

		switch route.direction {
		case north:
			y--
		case south:
			y++
		default:
			panic("unexpected direction")
		}

	}
	// TODO: found exit
	route.position = previousPosition

	return true, route
}

// Calculate the travel in East-West direction. Returns true if the route exist the grid.
// Otherwise false
func (route *guardRoute) travelEastWest(labMap []input.Coordinates) (bool, *guardRoute) {
	var trail []input.Coordinates

	previousPosition := route.position
	for x := route.position.X; x < route.mapSize.X; {
		route.position.X = x

		for _, block := range labMap {
			if input.SameCoordinates(block, route.position) {
				fmt.Printf("Found a block at position %v\n", previousPosition)
				route.position = previousPosition
				return false, route
			}
		}
		trail = append(trail, route.position) // TODO: this must be a set or something
		previousPosition = route.position

		switch route.direction {
		case east:
			x++
		case west:
			x--
		default:
			panic("unexpected direction")
		}
	}

	route.position = previousPosition
	return true, route
}

// Calculates the travel to the direction defined in the route structure. Returns updated route
// struck and boolean true if an exit was found. The boolean will be false otherwise
func (route *guardRoute) travel(labMap []input.Coordinates) (bool, *guardRoute) {
	debug := false
	var exitFound bool
	switch route.direction {
	case north:
		exitFound, route = route.travelNorthSouth(labMap)
	case east:
		exitFound, route = route.travelEastWest(labMap)
	case south:
		exitFound, route = route.travelNorthSouth(labMap)
	case west:
		exitFound, route = route.travelEastWest(labMap)
	default:
		panic("unexpected direction")
	}

	if debug {
		fmt.Println(labMap)
	}

	return exitFound, route
}

func (route *guardRoute) guardNavigation(blockMap []input.Coordinates) {
	var exitFound bool
	killSwitch := 1200
	i := 0
	j := 0
	allDirections := [...]Direction{north, east, south, west}

	for true {
		if j > killSwitch {
			panic("The kill switch was triggered")
		}

		route.direction = allDirections[i]
		exitFound, route = route.travel(blockMap)
		if exitFound {
			fmt.Printf("The guard will exit here %v\n", route.position)
			return
		}
		i++
		j++

		if i >= len(allDirections) {
			i = 0
		}

	}
}

type guardRoute struct {
	mapSize   input.Coordinates
	position  input.Coordinates
	trail     map[string]bool
	direction Direction
}

func CalculateRoute() {
	total := 0

	var route guardRoute
	inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	// inputLines := input.GetInputV2("../downloads/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")

	route.mapSize = input.GetGridSize(inputLines)
	route.position = startingPoint[0]
	route.direction = north

	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v guards in total.\n", len(startingPoint))
	fmt.Printf("The total size of the map is %v by %v\n", route.mapSize.X, route.mapSize.Y)

	route.guardNavigation(blockCoordinates)
	fmt.Println("-----")

	/*
		position := travel(route.mapSize, blockCoordinates, startingPoint[0], north)
		position = travel(route.mapSize, blockCoordinates, position, east)
		position = travel(route.mapSize, blockCoordinates, position, south)
		position = travel(route.mapSize, blockCoordinates, position, west)
	*/
	fmt.Printf("The lab guard will visit %v distinct positions\n", total)
}
