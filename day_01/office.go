package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
func parseTwoLists(inputList []string) ([]int, []int) {
	var leftInput []int
	var rightInput []int

	debug := false

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

func main() {
	test := false

	fmt.Println("Running")
	if test {
		testArrays()
		fmt.Println("-----")
	}
	inputLines := getInput("test-input-2.txt")
	fmt.Printf("i_lines type %T\n", inputLines)

	leftInput, rightInput := parseTwoLists(inputLines)

	for i, v := range leftInput {
		fmt.Printf("line %v:         %v\n", i, v)
	}

	for _, v := range rightInput {
		fmt.Printf("value: %v\n", v)
	}
}
