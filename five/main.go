package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Computer ...
type Computer struct {
	pointer      int
	instructions []int
}

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

func determineOperationAndModes(code int) (op int, modes []int) {
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

func normalizeModes(modes []int, lenParams int) []int {
	fixedModes := make([]int, lenParams)
	copy(fixedModes, modes)

	return fixedModes
}

func (c *Computer) initializeProgram(instructions []int) {
	// we create a copy of the instructions to prevent mutating the original
	ownedInstructions := make([]int, len(instructions))
	copy(ownedInstructions, instructions)

	c.pointer = 0
	c.instructions = ownedInstructions
}

func (c *Computer) nextPointer(value int) {
	c.pointer = value
}

func (c *Computer) resolveValues(params, modes []int) []int {
	modes = normalizeModes(modes, len(params))
	values := make([]int, len(params))

	for i, param := range params {
		switch modes[i] {
		case 0:
			values[i] = c.instructions[param]
		default:
			values[i] = param
		}
	}

	return values
}

func (c *Computer) add() {
	defer c.nextPointer(c.pointer + 4)
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+4]
	values := c.resolveValues(params, modes)
	c.instructions[params[2]] = values[0] + values[1]
}

func (c *Computer) mul() {
	defer c.nextPointer(c.pointer + 4)
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+4]
	values := c.resolveValues(params, modes)
	c.instructions[params[2]] = values[0] * values[1]
}

func (c *Computer) set(input int) {
	defer c.nextPointer(c.pointer + 2)

	params := c.instructions[c.pointer+1 : c.pointer+2]
	c.instructions[params[0]] = input
}

func (c *Computer) get() int {
	defer c.nextPointer(c.pointer + 2)
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+2]
	values := c.resolveValues(params, modes)

	return values[0]
}

func (c *Computer) jumpIfTrue() {
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+3]
	values := c.resolveValues(params, modes)

	switch {
	case values[0] != 0:
		c.nextPointer(values[1])
	default:
		c.nextPointer(c.pointer + 3)
	}
}

func (c *Computer) jumpIfFalse() {
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+3]
	values := c.resolveValues(params, modes)

	switch {
	case values[0] == 0:
		c.nextPointer(values[1])
	default:
		c.nextPointer(c.pointer + 3)
	}
}

func (c *Computer) lessThan() {
	defer c.nextPointer(c.pointer + 4)
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+4]
	values := c.resolveValues(params, modes)
	switch values[0] < values[1] {
	case true:
		c.instructions[params[2]] = 1
	default:
		c.instructions[params[2]] = 0
	}
}

func (c *Computer) equals() {
	defer c.nextPointer(c.pointer + 4)
	_, modes := determineOperationAndModes(c.instructions[c.pointer])

	params := c.instructions[c.pointer+1 : c.pointer+4]
	values := c.resolveValues(params, modes)
	switch values[0] == values[1] {
	case true:
		c.instructions[params[2]] = 1
	default:
		c.instructions[params[2]] = 0
	}
}

func (c *Computer) runProgram(instructions []int, arg int) []int {
	c.initializeProgram(instructions)
	log := make([]int, 0)

program:
	for {
		opCode, _ := determineOperationAndModes(c.instructions[c.pointer])

		switch opCode {
		case 0:
			c.nextPointer(c.pointer + 1)
		case 1:
			c.add()
		case 2:
			c.mul()
		case 3:
			c.set(arg)
		case 4:
			log = append(log, c.get())
		case 5:
			c.jumpIfTrue()
		case 6:
			c.jumpIfFalse()
		case 7:
			c.lessThan()
		case 8:
			c.equals()
		case 99:
			break program
		}
	}

	return log
}

func main() {
	instructions := parseInstructions("./instructions.txt")
	computer := new(Computer)
	log := computer.runProgram(instructions, 1)

	// this is the answer to part one
	fmt.Println(log)

	log2 := computer.runProgram(instructions, 5)

	// this is the answer to part two
	fmt.Println(log2)
}
