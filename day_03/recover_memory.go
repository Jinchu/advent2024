package day_03

import (
	"advent/internal/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseCommands(memoryFragment string) []string {
	var commands []string
	regex_mask := regexp.MustCompile("^\\(\\d{1,3},\\d{1,3}\\)")

	potential_commands := strings.Split(memoryFragment, "mul")

	for _, candiate := range potential_commands {
		if regex_mask.MatchString(candiate) {
			commands = append(commands, candiate)
		}
	}

	return commands
}

func parseConditonalStatements(memoryFragment string) []string {
	var conditionals []string

	dirtyConditional := strings.Split(memoryFragment, "do()")
	for _, dirty := range dirtyConditional {
		var cleaned string
		cleaned = strings.Split(dirty, "don't()")[0]
		conditionals = append(conditionals, cleaned)
	}

	return conditionals
}

// computeCommand() calculates the multiplication in the command.
func computeCommand(command string) int {
	debug := false

	sides := strings.Split(command, ",")
	left_str := strings.Split(sides[0], "(")[1]
	right_str := strings.Split(sides[1], ")")[0]

	left, _ := strconv.Atoi(left_str)
	right, _ := strconv.Atoi(right_str)
	if debug {
		fmt.Println(command)
		fmt.Printf("%v * %v\n", left, right)
		fmt.Println("----------")
	}

	return left * right
}

/*
...
7 * 32
*/
func RecoverMemory() {
	inputLines := input.GetInputV2("./day_03/input-day3.txt")
	// inputLines := input.GetInputV2("./day_03/test-input-2.txt")
	total := 0
	singleLine := ""
	for _, mFragment := range inputLines {
		lineContent := strings.TrimSpace(string(mFragment))
		singleLine = singleLine + lineContent
	}
	cleanedFragments := parseConditonalStatements(singleLine)
	for _, fragment := range cleanedFragments {
		commands := parseCommands(fragment)
		for _, c := range commands {
			singleResult := computeCommand(c)
			total = total + singleResult
		}
	}

	fmt.Printf("Result: %v\n\n", total)
}
