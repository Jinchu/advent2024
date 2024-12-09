package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type inputs []string

// is this a comment
func get_input(filep string) inputs {
	f, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}

	// defer so that we close the file always
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input_lines inputs
	for scanner.Scan() {
		input_lines = append(input_lines, scanner.Text())
	}

	return input_lines
}

func main() {
	fmt.Println("Hello, World!")
	i_lines := get_input("test-input-2.txt")
	fmt.Printf("i_lines type %T\n", i_lines)
	// 	fmt.Println(i_lines)

	for i, v := range i_lines {

		fmt.Printf("line %v:         %v\n", i, v)
	}
}
