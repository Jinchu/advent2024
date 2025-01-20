package day06

import (
	"advent/internal/input"
	"fmt"
	"strconv"
	"strings"
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
	previousPosition := route.position
	for y := route.position.Y; y < route.mapSize.Y; {

		route.position.Y = y
		for _, block := range labMap {
			if input.SameCoordinates(block, route.position) {
				// Hitting a block
				route.position = previousPosition
				return false, route
			}
		}
		trailPoint := convertToStr(route.position)
		route.trail[trailPoint] = true
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
	route.position = previousPosition
	return true, route // Found the exit
}

// Calculate the travel in East-West direction. Returns true if the route exist the grid.
// Otherwise false
func (route *guardRoute) travelEastWest(labMap []input.Coordinates) (bool, *guardRoute) {
	previousPosition := route.position
	for x := route.position.X; x < route.mapSize.X; {
		route.position.X = x

		for _, block := range labMap {
			if input.SameCoordinates(block, route.position) {
				route.position = previousPosition
				return false, route
			}
		}
		trailPoint := convertToStr(route.position)
		route.trail[trailPoint] = true
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
	return true, route // Found the exit
}

// addPading prep ends the ordinal string with zeroes so that the length is always 3
func addPading(orginalString string) string {
	var padded string
	switch len(orginalString) {
	case 1:
		padded = "00" + orginalString
	case 2:
		padded = "0" + orginalString
	case 3:
		padded = orginalString
	default:
		panic("Cannot convert to sting with these assumptions")
	}

	return padded
}

func convertToStr(originalObject input.Coordinates) string {
	var coordinateStr string

	xStr := addPading(strconv.Itoa(originalObject.X))
	yStr := addPading(strconv.Itoa(originalObject.Y))

	coordinateStr = xStr + "," + yStr

	return coordinateStr
}

func convertToCoordinate(coordinateString string) input.Coordinates {
	var converted input.Coordinates

	splitted := strings.Split(coordinateString, ",")

	xInt, err := strconv.Atoi(splitted[0])
	converted.X = xInt
	if err != nil {
		errorStr := fmt.Sprintf("Cannot convert %v", coordinateString)
		panic(errorStr)
	}

	yInt, err := strconv.Atoi(splitted[1])
	converted.Y = yInt
	if err != nil {
		errorStr := fmt.Sprintf("Cannot convert %v", coordinateString)
		panic(errorStr)
	}

	return converted
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
	var route guardRoute
	inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	// inputLines := input.GetInputV2("../downloads/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")

	route.mapSize = input.GetGridSize(inputLines)
	route.position = startingPoint[0]
	route.direction = north
	route.trail = make(map[string]bool)

	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v guards in total.\n", len(startingPoint))
	fmt.Printf("The total size of the map is %v by %v\n", route.mapSize.X, route.mapSize.Y)

	route.guardNavigation(blockCoordinates)
	fmt.Println("-----")

	if len(route.trail) != 41 {
		panic("You have broken the thing")
	}

	for i, v := range route.trail {
		fmt.Printf("%v - %v\n", i, v)
	}

	quickTest := convertToCoordinate("005,008")
	fmt.Println(quickTest)

	fmt.Printf("The lab guard will visit %v distinct positions\n", len(route.trail))
}
