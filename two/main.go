package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(bytes)
	textArray := strings.Split(text, ",")

	intArray := make([]int, len(textArray))
	for i, value := range textArray {
		parsed, _ := strconv.ParseInt(value, 10, 64)
		intArray[i] = int(parsed)
	}

	return intArray
}

func correctInput(input []int) {
	input[1] = 12
	input[2] = 2
}

func runProgram(input []int) {
	cursor := 0

	for input[cursor] != 99 {
		posA := input[cursor+1]
		posB := input[cursor+2]
		posC := input[cursor+3]

		switch input[cursor] {
		case 1:
			input[posC] = input[posA] + input[posB]
		case 2:
			input[posC] = input[posA] * input[posB]
		}

		cursor += 4
	}
}

func main() {
	input := parseInput("./input.txt")
	correctInput(input)

	runProgram(input)

	// this is the answer to part 1
	fmt.Println(input[0])
}
