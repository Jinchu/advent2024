package main

import (
	"fmt"
	"office/internal/input"
	"sort"
)

func calculateDistance(leftSlice []int, rightSlice []int) uint {
	var cumulativeDistance uint = 0

	for i, v := range leftSlice {
		currentDistance := v - rightSlice[i]
		if currentDistance < 0 {
			currentDistance = 0 - currentDistance
		}
		cumulativeDistance = cumulativeDistance + uint(currentDistance)
		// fmt.Printf("Current: %v  Total: %v\n", currentDistance, cumulativeDistance)
	}

	return cumulativeDistance
}

func main() {
	debug := false
	test := false
	fmt.Println("Running")
	if test {
		testArrays()
		fmt.Println("-----")
	}
	inputLines := input.GetInput("input-day1.txt")

	leftInput, rightInput := input.ParseTwoLists(inputLines, debug)

	sort.Slice(leftInput, func(i, j int) bool { return leftInput[i] < leftInput[j] })
	sort.Slice(rightInput, func(i, j int) bool { return rightInput[i] < rightInput[j] })

	if debug {
		for i, v := range leftInput {
			fmt.Printf("%v  %v\n", v, rightInput[i])
		}
	}

	totalDistance := calculateDistance(leftInput, rightInput)
	fmt.Println(totalDistance)
}
