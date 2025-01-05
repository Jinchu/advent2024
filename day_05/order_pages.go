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
	debug := false
	// inputLines := input.GetInputV2("./day_05/test-input-1.txt")
	//	inputLines := input.GetInputV2("./day_05/input-day5.txt")

	rules, pageNumbers := getRulesAndPages("./day_05/test-input-1.txt")

	if debug {
		for _, update := range pageNumbers {
			fmt.Println(update[2])
			if rules[2].before == 97 {
				fmt.Println("dummy")
			}
		}
	}

	for _, update := range pageNumbers {
		isValid := checkUpdate(update, rules)
		if isValid {
			fmt.Printf("Valid update found %v\n", update)
		}
	}
}

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
			fmt.Printf("Left check: line %v breaks the rule %v\n", wholeUpdate, currentRule)
			fmt.Printf("Current page number [%v] %v\n\n", checkIndex, checkNumber)
			return false
		}
	}

	return true // no rule braking numbers found

}

func checkRight(wholeUpdate []int, currentRule rule, checkIndex int) bool {
	if currentRule.after != wholeUpdate[checkIndex] {
		return true // rule is not applicable
	}

	for i := checkIndex + 1; i < len(wholeUpdate); i++ {
		var checkNumber int
		checkNumber = wholeUpdate[i]

		if checkNumber == currentRule.before {
			fmt.Printf("Right check: line %v breaks the rule %v\n", wholeUpdate, currentRule)
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
