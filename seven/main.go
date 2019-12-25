package main

import (
	"fmt"
	"sort"
)

func valueExists(values []int, value int) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}

	return false
}

func copyValues(values []int) []int {
	copied := make([]int, len(values), cap(values))
	copy(copied, values)

	return copied
}

func _getPermutations(chosenValues []int, values []int, repeat bool) [][]int {
	permutations := make([][]int, 0)

	for _, val := range values {
		if repeat == false && valueExists(chosenValues, val) {
			continue
		}

		nextChosenValues := copyValues(chosenValues)
		nextChosenValues = append(nextChosenValues, val)
		if len(nextChosenValues) == cap(nextChosenValues) {
			permutations = append(permutations, nextChosenValues)
			continue
		}

		permutations = append(permutations, _getPermutations(nextChosenValues, values, repeat)...)
	}

	return permutations
}

func getPermutations(values []int, k int, repeat bool) [][]int {
	return _getPermutations(make([]int, 0, k), values, repeat)
}

func (c *Computer) runAmplifier(instructions []int, phase, inputSignal int) int {
	log := c.runProgram(instructions, []int{phase, inputSignal})
	return log[0]
}

func main() {
	instructions := parseInstructions("./instructions.txt")
	c := new(Computer)

	phases := []int{0, 1, 2, 3, 4}
	phasePermutations := getPermutations(phases, len(phases), false)

	outputs := make([]int, 0)
	for _, phasePerm := range phasePermutations {
		nextInputSignal := 0
		for _, phase := range phasePerm {
			nextInputSignal = c.runAmplifier(instructions, phase, nextInputSignal)
		}
		outputs = append(outputs, nextInputSignal)
	}

	sort.Ints(outputs)

	// this is the answer to part 1
	fmt.Println(outputs[len(outputs)-1])
}
