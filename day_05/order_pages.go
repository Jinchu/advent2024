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
	// inputLines := input.GetInputV2("./day_05/test-input-1.txt")
	//	inputLines := input.GetInputV2("./day_05/input-day5.txt")

	rules, pageNumbers := getRulesAndPages("./day_05/test-input-1.txt")
	for _, update := range pageNumbers {
		fmt.Println(update[2])
		if rules[2].before == 97 {
			fmt.Println("dummy")
		}
	}
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
				fmt.Printf("Cannot convert %v to integer", parts[0])
				panic(err)
			}

			afterInt, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Cannot convert %v to integer", parts[1])
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
