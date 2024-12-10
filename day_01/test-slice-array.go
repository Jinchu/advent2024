package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type inputs_arr [8]string

// is this a comment
func getInputTT(filep string) []string {
	f, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}

	// defer so that we close the file always
	defer f.Close()
	var inputLines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
		fmt.Printf("cap %v, len %v, %p\n", cap(inputLines), len(inputLines), inputLines)
	}

	return inputLines
}

func getInputToArray(inputs_arr *inputs_arr, filep string) {
	f, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}

	// defer so that we close the file always
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {

		inputs_arr[i] = scanner.Text()
		i++
	}

	return
}

func testArrays() {
	fmt.Println("Hello, World!")
	var inputArray inputs_arr
	input_lines := getInputTT("test-input-2.txt")
	fmt.Printf("i_lines type %T\n", input_lines)
	// 	fmt.Println(i_lines)

	getInputToArray(&inputArray, "test-input-2.txt")

	for i, v := range input_lines {

		fmt.Printf("line %v:         %v\n", i, v)
	}

	fmt.Println("----")
	for _, v := range inputArray {
		fmt.Printf("line: %v\n", v)
	}
}
