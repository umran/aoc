package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInstructions(filename string) []int {
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

func determineOperation(code int) (op int, modes []int) {
	codeString := strconv.FormatInt(int64(code), 10)
	slices := strings.Split(codeString, "")

	// start reading the opcode digits in reverse
	for i := len(slices) - 1; i >= 0; i-- {
		value, _ := strconv.ParseInt(slices[i], 10, 64)
		switch i {
		case len(slices) - 1:
			op += int(value)
		case len(slices) - 2:
			op += int(value) * 10
		default:
			modes = append(modes, int(value))
		}
	}

	return op, modes
}

func runProgram(instructions []int, input int) []int {
	pointer := 0
	log := make([]int, 0)

program:
	for {
		opCode, modes := determineOperation(instructions[pointer])

		switch opCode {
		case 0:
			pointer++
		case 1:
			fmt.Println("detected an addition operation")
			fixedModes := make([]int, 3)
			copy(fixedModes, modes)

			var (
				a int
				b int
			)

			switch fixedModes[0] {
			case 0:
				a = instructions[instructions[pointer+1]]
			default:
				a = instructions[pointer+1]
			}

			switch fixedModes[1] {
			case 0:
				b = instructions[instructions[pointer+2]]
			default:
				b = instructions[pointer+2]
			}

			result := a + b
			instructions[instructions[pointer+3]] = result

			pointer += 4
		case 2:
			fmt.Println("detected a multiplication operation")
			fixedModes := make([]int, 3)
			copy(fixedModes, modes)

			var (
				a int
				b int
			)

			switch fixedModes[0] {
			case 0:
				a = instructions[instructions[pointer+1]]
			default:
				a = instructions[pointer+1]
			}

			switch fixedModes[1] {
			case 0:
				b = instructions[instructions[pointer+2]]
			default:
				b = instructions[pointer+2]
			}

			result := a * b
			instructions[instructions[pointer+3]] = result

			pointer += 4
		case 3:
			fmt.Println("detected a set operation")
			// fixedModes := make([]int, 1)
			// copy(fixedModes, modes)
			instructions[instructions[pointer+1]] = input

			pointer += 2
		case 4:
			fmt.Println("detected a get operation")
			fixedModes := make([]int, 1)
			copy(fixedModes, modes)

			var getPointer int
			switch fixedModes[0] {
			case 0:
				getPointer = instructions[pointer+1]
			default:
				getPointer = pointer + 1
			}

			log = append(log, instructions[getPointer])

			pointer += 2
		case 99:
			fmt.Println("detected a halt instruction")
			break program
		}
	}

	return log
}

func main() {
	instructions := parseInstructions("./instructions.txt")
	log := runProgram(instructions, 1)
	fmt.Println(log)
}
