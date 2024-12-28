package day_04

import (
	"advent/internal/input"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

func getVerticalLines(horizontal []string) []string {
	var verticalLines []string
	lineLenght := utf8.RuneCountInString(horizontal[0])

	for i := 0; i < lineLenght; i++ {

		currentVertical := ""
		for _, currentHorisontal := range horizontal {
			currentChar := string([]rune(currentHorisontal)[i])
			currentVertical = currentVertical + currentChar
		}
		verticalLines = append(verticalLines, currentVertical)
	}

	return verticalLines
}

func getDiagonalLines(horizontal []string) []string {
	var diagonalLines []string

	lineLenght := utf8.RuneCountInString(horizontal[0])
	colunmLenght := len(horizontal)
	var maxLen int
	if lineLenght > colunmLenght {
		maxLen = lineLenght
	} else {
		maxLen = colunmLenght
	}

	for i := 0; i < maxLen; i++ {
		x := i
		y := 0

		currentDiagonal := ""

		for j := y; j <= i; j++ {
			checkLine := horizontal[x]
			currentChar := string([]rune(checkLine)[y])
			currentDiagonal = currentDiagonal + currentChar
			x--
			y++
		}

		diagonalLines = append(diagonalLines, currentDiagonal)

	}

	diagonalLines = diagonalLines[:len(diagonalLines)-1]

	for i := 0; i <= maxLen; i++ {
		x := lineLenght - 1
		y := i

		currentDiagonal := ""

		for j := y; j < colunmLenght; j++ {
			checkLine := horizontal[x]
			currentChar := string([]rune(checkLine)[y])
			currentDiagonal = currentDiagonal + currentChar
			x--
			y++
		}

		diagonalLines = append(diagonalLines, currentDiagonal)
	}

	return diagonalLines
}

func printMatrix(outLines []string) {
	for _, line := range outLines {
		fmt.Println(line)
	}
}

func reverseString(orig string) string {
	runes := []rune(orig)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countMatches(direction []string, total int) int {
	searchWord := "XMAS"
	for _, line := range direction {
		total = total + strings.Count(line, searchWord)
		reversed := reverseString(line)
		total = total + strings.Count(reversed, searchWord)
	}

	return total
}

func checkAdjacent(grid []string, x int, y int) bool {
	checkLine := grid[y-1]
	upperLeft := string([]rune(checkLine)[x-1])
	upperRight := string([]rune(checkLine)[x+1])

	checkLine = grid[y+1]
	lowerLeft := string([]rune(checkLine)[x-1])
	lowerRight := string([]rune(checkLine)[x+1])

	if upperLeft == string('M') && lowerRight == string('S') {
		if upperRight == string('M') && lowerLeft == string('S') {
			return true
		} else if upperRight == string('S') && lowerLeft == string('M') {
			return true
		}
	} else if upperLeft == string('S') && lowerRight == string('M') {
		if upperRight == string('M') && lowerLeft == string('S') {
			return true
		} else if upperRight == string('S') && lowerLeft == string('M') {
			return true
		}
	}

	return false
}

func gridTravelCount(inputLines []string) uint {
	var totalFound uint
	firstLine := inputLines[0]

	totalFound = 0
	for y := 1; y < len(inputLines)-1; y++ {
		for x := 1; x < len(firstLine)-1; x++ {
			currentLine := inputLines[y]
			currentChar := string([]rune(currentLine)[x])
			if currentChar == string('A') {

				if checkAdjacent(inputLines, x, y) {
					totalFound++
				}
			}
		}
	}

	return totalFound
}

// part 2
func SearchForX() {
	// inputLines := input.GetInputV2("./day_04/test-input-1.txt")
	inputLines := input.GetInputV2("./day_04/input-day4.txt")

	matchesFound := gridTravelCount(inputLines)

	fmt.Printf("Found %v matches\n\n", matchesFound)
}

// part 1
func SearchForChristmas() {
	total := 0

	inputLines := input.GetInputV2("./day_04/test-input-1.txt")
	// inputLines := input.GetInputV2("./day_04/input-day4.txt")
	reverseInput := slices.Clone(inputLines)
	slices.Reverse(reverseInput)

	verticalLines := getVerticalLines(inputLines)
	diagonalLines := getDiagonalLines(inputLines)
	secondDiag := getDiagonalLines(reverseInput)

	total = countMatches(inputLines, total)
	total = countMatches(verticalLines, total)
	total = countMatches(diagonalLines, total)
	total = countMatches(secondDiag, total)

	fmt.Println("\n---")
	fmt.Printf("Found the word %v times\n", total)
}
