package day_04

import (
	"advent/internal/input"
	"fmt"
	"unicode/utf8"
)

func getVerticalLines(horisontal []string) []string {

	var verticalLines []string
	lineLeght := utf8.RuneCountInString(horisontal[0])

	for i := 0; i < lineLeght; i++ {

		currentVertical := ""
		for _, currentHorisontal := range horisontal {
			currentChar := string([]rune(currentHorisontal)[i])
			currentVertical = currentVertical + currentChar
		}
		verticalLines = append(verticalLines, currentVertical)
	}

	return verticalLines
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
	printMatrix(verticalLines)

	fmt.Println("\n---")
	fmt.Printf("Found the word %v times\n", total)
}
