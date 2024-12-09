package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type inputs []string
type inputs_arr [8]string

// is this a comment
func get_input(input_lines inputs, filep string) inputs {
	f, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}

	// defer so that we close the file always
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input_lines = append(input_lines, scanner.Text())
	}

	return input_lines
}

func get_input_to_array(inputs_arr *inputs_arr, filep string) {
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

func main() {
	fmt.Println("Hello, World!")
	var input_lines inputs
	var input_array inputs_arr
	input_lines = get_input(input_lines, "test-input-2.txt")
	fmt.Printf("i_lines type %T\n", input_lines)
	// 	fmt.Println(i_lines)

	get_input_to_array(&input_array, "test-input-2.txt")

	for i, v := range input_lines {

		fmt.Printf("line %v:         %v\n", i, v)
	}

	fmt.Println("----")
	for _, v := range input_array {
		fmt.Printf("line: %v\n", v)
	}
}
