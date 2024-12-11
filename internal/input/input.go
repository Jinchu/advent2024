package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetInput reads the input file and returns a slice containing the rows of the file.
func GetInput(filep string) []string {
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

// ParsesIntFromString parses list of unsigned integers from a given string
func ParseIntFromStrint(inputString string, debug bool) []uint {
	c := strings.Split(inputString, " ")
	var report []uint
	for _, letter := range c {
		level_int, err := strconv.Atoi(letter)
		if err != nil {
			// This is fine. Not all letters are meaningful
		} else {
			report = append(report, uint(level_int))
		}
	}

	if debug {
		fmt.Printf("Returning a report with the length of %v\n\n", len(report))
	}

	return report
}

// ParseTwoLists returns two slices containing the two lists in the input
func ParseTwoLists(inputList []string, debug bool) ([]int, []int) {
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
