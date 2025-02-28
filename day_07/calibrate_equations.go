package day07

import (
	"advent/internal/input"
	"fmt"
	"strconv"
	"strings"
)

/* getTestResultCalibrationNumbers() converts the string to test number and calibration numbers.
 */
func getTestResultCalibrationNumbers(cLine string) (int, []int) {
	var calibrationNumbers []int

	sides := strings.Split(cLine, ":")
	testResult, err := strconv.Atoi(sides[0])
	if err != nil {
		panic("Cant convert input to int")
	}

	rightStrings := strings.Split(sides[1], " ")
	for i, numberStr := range rightStrings {
		if i == 0 {
			continue
		}
		currentNum, _ := strconv.Atoi(numberStr)
		calibrationNumbers = append(calibrationNumbers, currentNum)
	}

	return testResult, calibrationNumbers
}

func concatTwoIntegersAsStr(left int, right int) int {
	leftStr := strconv.Itoa(left)
	rightStr := strconv.Itoa(right)
	totalStr := leftStr + rightStr

	totalInt, _ := strconv.Atoi(totalStr)

	// fmt.Printf("%v || %v - %v\n", left, right, totalInt)
	return totalInt
}

func calculateNextNumbers(cumulativeLeft []int, right int) []int {
	var possibilities []int

	for _, left := range cumulativeLeft {
		sum := left + right
		product := left * right
		concat := concatTwoIntegersAsStr(left, right)

		possibilities = append(possibilities, sum)
		possibilities = append(possibilities, product)
		possibilities = append(possibilities, concat)
	}

	return possibilities
}

func getPossibleCalibrationResults(calibrationNumbers []int) []int {
	var cumulativeLeft []int

	for i, currentNum := range calibrationNumbers {
		if i == 0 {
			cumulativeLeft = append(cumulativeLeft, currentNum)
			continue
		}

		cumulativeLeft = calculateNextNumbers(cumulativeLeft, currentNum)

	}
	return cumulativeLeft

}

/* getLineCalibrationValue return the line specific calibration value. Returns 0 for invalid lines.
 */
func getLineCalibrationValue(cLine string) int {
	total := 0

	testResult, calibrationNumbers := getTestResultCalibrationNumbers(cLine)

	// test operations
	possibleResults := getPossibleCalibrationResults(calibrationNumbers)
	// fmt.Println(possibleResults)

	for _, possible := range possibleResults {
		if possible == testResult {
			total = testResult
		}
	}

	// fmt.Printf("line total: %v\n", total)

	return total
}

func GetTotalCalibration() {
	// inputLines := input.GetInputV2("./day_07/test-input-1.txt")
	inputLines := input.GetInputV2("../downloads/input-day7.txt")
	total := 0

	for _, line := range inputLines {
		total = total + getLineCalibrationValue(line)
	}

	fmt.Printf("The total calibration value is: %v\n", total)
}
