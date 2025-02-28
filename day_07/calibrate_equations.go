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

/* getLineCalibrationValue return the line specific calibration value. Returns 0 for invalid lines.
 */
func getLineCalibrationValue(cLine string) int {
	total := 0

	testResult, calibrationNumbers := getTestResultCalibrationNumbers(cLine)
	fmt.Println(calibrationNumbers)

	fmt.Printf("Trying to get the total of %v: %v\n", testResult, total)
	return total
}

func GetTotalCalibration() {
	inputLines := input.GetInputV2("./day_07/test-input-1.txt")
	total := 0

	for _, line := range inputLines {
		total = total + getLineCalibrationValue(line)
	}

	fmt.Printf("The total calibration value is: %v\n", total)
}
