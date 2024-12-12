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

// isSafe checks the given report WITHOUT the Problem Dampener
func isSafe(report []uint) bool {

	if !isIncreasing(report) && !isDecreasing(report) {
		return false
	}

	if !safeDistance(report) {
		return false
	}

	return true
}

func remove(slice []uint, s int) []uint {
	return append(slice[:s], slice[s+1:]...)
}

// / isSafeWithDampener removes one level at the time and checks if the dampened report is safe
func isSafeWithDampener(report []uint) bool {
	original_report := make([]uint, len(report))
	copy(original_report, report)

	for i := range report {
		dampened := remove(report, i)
		if isSafe(dampened) {
			return true
		}
		copy(report, original_report)
	}
	return false
}

func CountSafeReports() {
	debug := false
	var safeCount uint = 0

	inputLines := input.GetInputV2("./day_02/input-day2.txt")
	// inputLines := input.GetInput("./day_02/test-input-1.txt")

	for _, rep := range inputLines {
		currentRep := input.ParseIntFromStrint(rep, debug)

		if isSafe(currentRep) {
			safeCount++
		} else if isSafeWithDampener(currentRep) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
	return
}
