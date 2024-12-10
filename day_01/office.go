package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// getInput reads the input file and returns a slice containing the rows of the file.
func getInput(filep string) []string {
	f, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}

	// defer so that we close the file always
	defer f.Close()
	var inputLines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	return inputLines
}

// parseTwoLists returns two slices containing the two lists in the input
func parseTwoLists(inputList []string, debug bool) ([]int, []int) {
	var leftInput []int
	var rightInput []int

	for _, v := range inputList {
		c := strings.Split(v, " ")
		if debug {
			fmt.Printf("c[0]: %v\n", c[0])
			fmt.Printf("c[3]: %v\n", c[3])
		}

		left, err := strconv.Atoi(c[0])
		if err != nil {
			fmt.Printf("Cannot convert %v to integer", c[0])
			panic(err)
		}
		right, err := strconv.Atoi(c[3])
		if err != nil {
			fmt.Printf("Cannot convert c[3] %v to integer\n", c[3])
			panic(err)
		}

		leftInput = append(leftInput, left)
		rightInput = append(rightInput, right)
	}

	return leftInput, rightInput
}

// printRawInput just prints every line of the input slice
func printRawInput(inputSlice []string) {
	for i, v := range inputSlice {
		fmt.Printf("line %v:         %v\n", i, v)
	}
}

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
	inputLines := getInput("input-day1.txt")

	leftInput, rightInput := parseTwoLists(inputLines, debug)

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
