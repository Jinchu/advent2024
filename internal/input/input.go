package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetInput reads the input file and returns a slice containing the rows of the file.
func GetInput(filep string) []string {
	f, _ := os.Open(filep)

	// defer so that we close the file always
	defer f.Close()
	var inputLines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	return inputLines
}

type Coordinates struct {
	X int
	Y int
}

// GetCoordinates parses all the coordiantes of the target string in the input. Returns a slice of
// struct coordinates.
func GetCoordinates(inputLines []string, target string) []Coordinates {
	firstLine := inputLines[0]
	var targetCoordinates []Coordinates

	for y := 0; y < len(inputLines); y++ {
		currentLine := inputLines[y]
		for x := 0; x < len(firstLine); x++ {
			currentChar := string([]rune(currentLine)[x])
			if currentChar == target {
				var newTarget Coordinates
				newTarget.X = x
				newTarget.Y = y
				targetCoordinates = append(targetCoordinates, newTarget)
			}
		}
	}

	return targetCoordinates
}

// Return the size of the gird as coordinates pointing to the maximum values.
func GetGridSize(inputLines []string) Coordinates {
	var size Coordinates
	size.Y = len(inputLines)
	size.X = len(inputLines[0])
	return size
}

// Returns true if the coordinates match
func SameCoordinates(comp1 Coordinates, comp2 Coordinates) bool {
	if comp1.X == comp2.X && comp1.Y == comp2.Y {
		return true
	} else {
		return false
	}
}

// GetInput reads the input file and returns a slice containing the rows of the file.
func GetInputV2(fileName string) []string {
	bytesRead, _ := os.ReadFile(fileName)
	fileContent := strings.TrimSpace(string(bytesRead))
	lines := strings.Split(fileContent, "\n")

	return lines
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
