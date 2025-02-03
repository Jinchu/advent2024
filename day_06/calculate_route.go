package day06

import (
	"advent/internal/input"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Direction int64

const (
	north Direction = iota
	east
	south
	west
)

var mutex sync.Mutex

// Calculate the travel in North-South direction. Returns true if the route exist the grid.
// Otherwise false
func (route *guardRoute) travelNorthSouth(labMap map[input.Coordinates]bool,
	loopDetection bool) (int64, *guardRoute) {
	previousPosition := route.position

	for y := route.position.Y; y < route.mapSize.Y; {
		if y < 0 {
			break
		}

		route.position.Y = y
		if labMap[route.position] {
			// Hitting a block
			route.position = previousPosition
			return 0, route
		}

		if loopDetection {
			trailPoint := convertToStrWDirection(route.position, route.direction)
			if route.trail[trailPoint] {
				// fmt.Printf("NS Loop detected %v\n", route.position)
				return 2, route
			}
			route.trail[trailPoint] = true
		} else {
			trailPoint := convertToStr(route.position)
			route.trail[trailPoint] = true
		}
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

	return 1, route // Found the exit
}

// Calculate the travel in East-West direction. Returns true if the route exist the grid.
// Otherwise false
func (route *guardRoute) travelEastWest(labMap map[input.Coordinates]bool,
	loopDetection bool) (int64, *guardRoute) {
	previousPosition := route.position
	for x := route.position.X; x < route.mapSize.X; {
		if x < 0 {
			break
		}
		route.position.X = x

		if labMap[route.position] {
			route.position = previousPosition
			return 0, route
		}

		if loopDetection {
			trailPoint := convertToStrWDirection(route.position, route.direction)
			if route.trail[trailPoint] {
				// fmt.Printf("EW Loop detected %v\n", route.position)
				return 2, route
			}
			route.trail[trailPoint] = true
		} else {
			trailPoint := convertToStr(route.position)
			route.trail[trailPoint] = true
		}
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
	return 1, route // Found the exit
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
		fmt.Printf("Original: %v\n", orginalString)
		panic("Cannot convert to sting with these assumptions")
	}

	return padded
}

// Returns a string representation of the input coordinates augmented with the given direction
func convertToStrWDirection(originalObject input.Coordinates, direction Direction) string {
	var coordinateStr string
	var directionStr string

	if originalObject.X < 0 || originalObject.Y < 0 {
		panic("Encountered illegal coordinates [with directions]")
	}
	xStr := addPading(strconv.Itoa(originalObject.X))
	yStr := addPading(strconv.Itoa(originalObject.Y))
	switch direction {
	case north:
		directionStr = "north"
	case east:
		directionStr = "east"
	case south:
		directionStr = "south"
	case west:
		directionStr = "west"
	}

	coordinateStr = xStr + "," + yStr + "," + directionStr

	return coordinateStr
}

// Returns a string representation of the input coordinates
func convertToStr(originalObject input.Coordinates) string {
	var coordinateStr string

	if originalObject.X < 0 || originalObject.Y < 0 {
		panic("Encountered illegal coordinates []")
	}
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
func (route *guardRoute) travel(labMap map[input.Coordinates]bool,
	loopDetection bool) (int64, *guardRoute) {
	debug := false
	var exitFound int64
	switch route.direction {
	case north:
		exitFound, route = route.travelNorthSouth(labMap, loopDetection)
	case east:
		exitFound, route = route.travelEastWest(labMap, loopDetection)
	case south:
		exitFound, route = route.travelNorthSouth(labMap, loopDetection)
	case west:
		exitFound, route = route.travelEastWest(labMap, loopDetection)
	default:
		panic("unexpected direction")
	}

	if debug {
		fmt.Println(labMap)
	}

	return exitFound, route
}

func (route *guardRoute) guardNavigation(blockMap map[input.Coordinates]bool, loopDetection bool) int {
	var exitFound int64
	killSwitch := 1200
	i := 0
	j := 0
	allDirections := [...]Direction{north, east, south, west}

	for true {
		if j > killSwitch {
			panic("The kill switch was triggered")
		}

		route.direction = allDirections[i]
		exitFound, route = route.travel(blockMap, loopDetection)
		if exitFound == 1 {
			return 0
		}
		if exitFound == 2 {
			// fmt.Printf("Loop detected at position %v\n", route.position)
			return 1
		}
		i++
		j++

		if i >= len(allDirections) {
			i = 0
		}

	}

	return 0
}

type guardRoute struct {
	mapSize   input.Coordinates
	position  input.Coordinates
	trail     map[string]bool
	direction Direction
}

func afterAddingNewBlock(wg *sync.WaitGroup, total *int, startingPoint input.Coordinates,
	blockCoordinates map[input.Coordinates]bool, currentAddition string,
	mapSize input.Coordinates) {
	var currenRoute guardRoute
	tryBlock := convertToCoordinate(currentAddition)

	currenRoute.trail = make(map[string]bool)
	currenRoute.position = startingPoint
	currenRoute.direction = north
	currenRoute.mapSize = mapSize

	// fmt.Printf("blockCoordinates %p\n", &blockCoordinates)
	improvedBlock := make(map[input.Coordinates]bool)

	for coordinate, v := range blockCoordinates {
		if v {
			improvedBlock[coordinate] = true
		}
	}

	improvedBlock[tryBlock] = true
	res := currenRoute.guardNavigation(improvedBlock, true)

	if res == 1 {
		// print(res)
		mutex.Lock()
		defer mutex.Unlock()
		*total++
	}

	wg.Done()

}

func CalculateRoute() {
	var wg sync.WaitGroup
	var route guardRoute
	// inputLines := input.GetInputV2("./day_06/test-input-1.txt")
	inputLines := input.GetInputV2("../downloads/input-day6.txt")
	blockCoordinates := input.GetCoordinates(inputLines, "#")
	startingPoint := input.GetCoordinates(inputLines, "^")

	blockMap := make(map[input.Coordinates]bool)
	for _, v := range blockCoordinates {
		blockMap[v] = true
	}

	route.mapSize = input.GetGridSize(inputLines)
	route.position = startingPoint[0]
	route.direction = north
	route.trail = make(map[string]bool)

	fmt.Printf("Found %v blocks in total.\n", len(blockCoordinates))
	fmt.Printf("Found %v guards in total.\n", len(startingPoint))
	fmt.Printf("The total size of the map is %v by %v\n", route.mapSize.X, route.mapSize.Y)

	route.guardNavigation(blockMap, false)
	fmt.Println("-----")
	fmt.Printf("lenght %v\n", len(route.trail))

	// Ignore the starting point
	total := 0
	startCoord := convertToStr(startingPoint[0])
	route.trail[startCoord] = false
	originalRoute := route.trail

	for coordinate, v := range originalRoute {
		if v {
			route.trail = make(map[string]bool)
			route.position = startingPoint[0]
			route.direction = north

			improvedMap := blockMap
			wg.Add(1)
			go afterAddingNewBlock(&wg, &total, startingPoint[0], improvedMap, coordinate,
				input.GetGridSize(inputLines))
		}
	}

	wg.Wait()

	fmt.Printf("The number of possible loops is: %v\n", total)
}
