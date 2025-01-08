package day_05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func PrintOrder() {
	total := 0
	debug := false

	// rules, pageNumbers := getRulesAndPages("./day_05/test-input-1.txt")
	rules, pageNumbers := getRulesAndPages("./day_05/input-day5.txt")

	for _, update := range pageNumbers {
		isValid := checkUpdate(update, rules)

		if debug {
			fmt.Println("---")
		}

		// Solution to part one can be found in branch day5-pt1
		if !isValid {
			if debug {
				fmt.Printf("Invalid update found %v\n", update)
			}
			fixedOrder := fixPrintOrder(update, rules)
			total = total + getMiddleValue(fixedOrder, debug)
		}
	}

	fmt.Printf("The sum of middle numbers from updates is: %v\n\n", total)
}

func getMiddleValue(update []int, debug bool) int {
	if len(update)%2 == 0 {
		fmt.Printf("ERROR: the length is an even number %v\n", len(update))
		fmt.Printf("%v\n", update)
		return 0
	}

	middle := (len(update) - 1) / 2

	if debug {
		fmt.Printf("the middle value is %v\n\n", update[middle])
	}

	return update[middle]
}

// checkUpdate return true if the update does not break any rules
func checkUpdate(wholeUpdate []int, ruleBook []rule) bool {
	for _, currentRule := range ruleBook {
		for i := range wholeUpdate {
			if !checkRight(wholeUpdate, currentRule, i) {
				return false
			}
			if !checkLeft(wholeUpdate, currentRule, i) {
				return false
			}
		}
	}
	return true
}

func checkLeft(wholeUpdate []int, currentRule rule, checkIndex int) bool {
	if currentRule.before != wholeUpdate[checkIndex] {
		return true // rule is not applicable
	}

	for i := checkIndex - 1; i > 0; i-- {
		var checkNumber int
		checkNumber = wholeUpdate[i]

		if checkNumber == currentRule.after {
			// fmt.Printf("Left check: line %v breaks the rule %v\n", wholeUpdate, currentRule)
			return false
		}
	}

	return true // no rule braking numbers found

}

func oneFixStep(update []int, ruleBook []rule) []int {
	var fixedUpdate []int
	for _, currentRule := range ruleBook {
		for i := range update {
			breakingIndex := checkRightWithIndex(update, currentRule, i)
			if breakingIndex < 0 {
				continue
			}
			fixedUpdate = conformToRule(update, i, breakingIndex)
		}
	}

	return fixedUpdate
}

func fixPrintOrder(wholeUpdate []int, ruleBook []rule) []int {
	circuitBreaker := 12
	fixedUpdate := oneFixStep(wholeUpdate, ruleBook)

	for !checkUpdate(fixedUpdate, ruleBook) {
		if circuitBreaker <= 0 {
			break
		}
		fixedUpdate = oneFixStep(fixedUpdate, ruleBook)
		circuitBreaker--
	}

	return fixedUpdate
}

func conformToRule(wholeUpdate []int, checkIndex int, breakingIndex int) []int {
	debug := false
	var fixedUpdate []int
	if breakingIndex < 0 {
		panic("Nothing to return")
	} else {
		if debug {
			fmt.Printf("The violating character is %v [position %v]\n",
				wholeUpdate[breakingIndex], breakingIndex)
			fmt.Printf("The violating character should be before %v [position %v]\n",
				wholeUpdate[checkIndex], checkIndex)
		}

		for i := range wholeUpdate {
			if i < checkIndex {
				fixedUpdate = append(fixedUpdate, wholeUpdate[i])
			} else if i > breakingIndex {
				fixedUpdate = append(fixedUpdate, wholeUpdate[i])
			} else if i == checkIndex {
				fixedUpdate = append(fixedUpdate, wholeUpdate[breakingIndex])
			} else if i == checkIndex+1 {
				fixedUpdate = append(fixedUpdate, wholeUpdate[checkIndex])
			}
		}

	}

	if debug {
		fmt.Printf("Fixed order: %v\n\n", fixedUpdate)
	}
	return fixedUpdate
}

// checkRight checks from left to right if the current rule is violated. Returns -1
// if a rule is not violated otherwise the index of the error
func checkRightWithIndex(wholeUpdate []int, currentRule rule, checkIndex int) int {
	if currentRule.after != wholeUpdate[checkIndex] {
		return -1 // rule is not applicable
	}

	for i := checkIndex + 1; i < len(wholeUpdate); i++ {
		var checkNumber int
		checkNumber = wholeUpdate[i]

		if checkNumber == currentRule.before {
			// fmt.Printf("Right check: line %v breaks the rule %v\n", wholeUpdate, currentRule)
			return i
		}
	}

	return -1 // no rule braking numbers found
}

func checkRight(wholeUpdate []int, currentRule rule, checkIndex int) bool {
	if currentRule.after != wholeUpdate[checkIndex] {
		return true // rule is not applicable
	}

	for i := checkIndex + 1; i < len(wholeUpdate); i++ {
		var checkNumber int
		checkNumber = wholeUpdate[i]

		if checkNumber == currentRule.before {
			// fmt.Printf("Right check: line %v breaks the rule %v\n", wholeUpdate, currentRule)
			return false
		}
	}

	return true // no rule braking numbers found
}

func convertUpdateNumbers(updateInput string) []int {
	var numbersInt []int

	rawNumbers := strings.Split(updateInput, ",")
	for _, num := range rawNumbers {
		cleanNumber, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		numbersInt = append(numbersInt, cleanNumber)
	}

	return numbersInt
}

func getRulesAndPages(fileName string) ([]rule, [][]int) {
	var rules []rule
	var pageNumbers [][]int
	bytesRead, _ := os.ReadFile(fileName)
	fileContent := strings.TrimSpace(string(bytesRead))
	lines := strings.Split(fileContent, "\n")

	for _, l := range lines {
		if strings.Contains(l, "|") {
			parts := strings.Split(l, "|")
			beforeInt, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("\nERROR: Cannot convert %v to integer\n", parts[0])
				panic(err)
			}

			afterInt, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("\nERROR: Cannot convert %v to integer\n", parts[1])
				panic(err)
			}

			newRule := rule{beforeInt, afterInt}
			rules = append(rules, newRule)
		} else if len(l) < 2 {
			continue
		} else {
			updateNumbers := convertUpdateNumbers(l)
			pageNumbers = append(pageNumbers, updateNumbers)
		}

	}

	return rules, pageNumbers
}
