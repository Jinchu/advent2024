package day_04

import (
	"advent/internal/input"
	"fmt"
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

func SearchForChristmas() {
	total := 0

	inputLines := input.GetInputV2("./day_04/test-input-1.txt")
	verticalLines := getVerticalLines(inputLines)

	printMatrix(inputLines)
	fmt.Println("")
	diagonalLines := getDiagonalLines(inputLines)
	printMatrix(diagonalLines)
	secondDiag := getDiagonalLines(verticalLines)
	printMatrix(secondDiag)

	fmt.Println("\n---")
	fmt.Printf("Found the word %v times\n", total)
}
