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

func copyInput(input []int) []int {
	newInput := make([]int, len(input))
	copy(newInput, input)

	return newInput
}

func adjustInput(input []int, noun, verb int) {
	input[1] = noun
	input[2] = verb
}

func runProgram(input []int) {
	pointer := 0

	for input[pointer] != 99 {
		valA := input[pointer+1]
		valB := input[pointer+2]
		valC := input[pointer+3]

		switch input[pointer] {
		case 1:
			input[valC] = input[valA] + input[valB]
		case 2:
			input[valC] = input[valA] * input[valB]
		}

		pointer += 4
	}
}

func findParams(input []int, condition int) (noun, verb int) {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			newInput := copyInput(input)
			adjustInput(newInput, i, j)
			runProgram(newInput)

			if newInput[0] == condition {
				noun = i
				verb = j
				break
			}
		}
	}

	return noun, verb
}

func main() {
	input := parseInput("./input.txt")

	newInput := copyInput(input)
	adjustInput(newInput, 12, 2)

	runProgram(newInput)

	// this is the answer to part 1
	fmt.Println(newInput[0])

	noun, verb := findParams(input, 19690720)

	// this is the answer to part 2
	fmt.Println(100*noun + verb)
}
