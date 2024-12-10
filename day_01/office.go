package main

import (
	"fmt"
	"office/internal/input"
)

// countMatches counts the number of times the selector is present in the targetSlice
func countMatches(selector int, targetSlice []int) int {
	matchCount := 0

	for _, v := range targetSlice {
		if selector == v {
			matchCount++
		}
	}
	return matchCount
}

func calculateSimilarity(leftInput []int, rightSlice []int) int {
	cumulativeSimilarityFactor := 0

	for _, leftV := range leftInput {
		matches := countMatches(leftV, rightSlice)
		cumulativeSimilarityFactor = cumulativeSimilarityFactor + (matches * leftV)

	}
	return cumulativeSimilarityFactor
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

	leftInput, rightInput := input.ParseTwoLists(inputLines, false)

	if debug {
		for i, v := range leftInput {
			fmt.Printf("%v  %v\n", v, rightInput[i])
		}
	}

	totalDistance := calculateSimilarity(leftInput, rightInput)
	fmt.Println(totalDistance)
}
