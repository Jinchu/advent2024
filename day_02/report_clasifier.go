package day_02

import (
	"advent/internal/input"
	"fmt"
)

func isDecreasing(report []uint) bool {
	isDecreasing := true
	for i, v := range report {
		if i == 0 {
			continue // skip the first object
		}

		if v >= report[i-1] {
			isDecreasing = false
		}
	}

	return isDecreasing
}

func isIncreasing(report []uint) bool {
	isIncreasing := true
	for i, v := range report {
		if i == 0 {
			continue // skip the first object
		}

		if v <= report[i-1] {
			isIncreasing = false
		}
	}

	return isIncreasing
}

func safeDistance(report []uint) bool {
	const safeTreshold = 3

	for i, v := range report {
		var distance int
		if i == 0 {
			continue // skip the first object
		}
		distance = int(v - report[i-1])
		if distance > 3 || distance < -3 {
			return false
		}
	}

	return true
}

func isSafe(report []uint) bool {
	isSafe := true

	if isIncreasing(report) || isDecreasing(report) {
		isSafe = true
	} else {
		isSafe = false
		return isSafe
	}

	if safeDistance(report) {
		isSafe = true
	} else {
		isSafe = false
	}

	return isSafe
}

func CountSafeReports() {
	debug := false
	var safeCount uint = 0

	inputLines := input.GetInput("./day_02/input-day2.txt")
	for _, rep := range inputLines {
		currentRep := input.ParseIntFromStrint(rep, debug)

		if isSafe(currentRep) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
	return
}
