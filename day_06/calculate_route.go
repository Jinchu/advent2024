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

func (route *guardRoute) travelNorthSouth(labMap []input.Coordinates) *guardRoute {
	var trail []input.Coordinates

	previousPosition := route.position
	for y := route.position.Y; y < route.mapSize.Y; {

		route.position.Y = y
		for _, block := range labMap {
			if sameCoordinates(block, route.position) {
				// Hitting a block
				fmt.Printf("Found a block at position %v\n", previousPosition)
				route.position = previousPosition
				return route
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
	fmt.Printf("The guard will exit here %v\n", route.position)

	return route
}

func (route *guardRoute) travelEastWest(labMap []input.Coordinates) *guardRoute {
	var trail []input.Coordinates

	previousPosition := route.position
	for x := route.position.X; x < route.mapSize.X; {
		route.position.X = x

		for _, block := range labMap {
			if sameCoordinates(block, route.position) {
				fmt.Printf("Found a block at position %v\n", previousPosition)
				route.position = previousPosition
				return route
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
	return route
}

func (route *guardRoute) travel(labMap []input.Coordinates) *guardRoute {
	debug := false
	switch route.direction {
	case north:
		route = route.travelNorthSouth(labMap)
	case east:
		route = route.travelEastWest(labMap)
	case south:
		route = route.travelNorthSouth(labMap)
	case west:
		route = route.travelEastWest(labMap)
	default:
		panic("unexpected direction")
	}

	if debug {
		fmt.Println(labMap)
	}

	return route
}

func (route *guardRoute) guardNavigation(blockMap []input.Coordinates) {
	i := 0
	j := 0
	allDirections := [...]Direction{north, east, south, west}

	for true {
		if j > 12 {
			panic("oops")
		}

		route.direction = allDirections[i]
		route = route.travel(blockMap)
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

	route.mapSize = getGridSize(inputLines)
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
