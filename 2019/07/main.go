package main

import (
	"fmt"
	"sort"
)

func main() {
	instructions := parseInstructions("./instructions.txt")

	// part 1
	phases := []int{0, 1, 2, 3, 4}
	phasePermutations := getPermutations(phases, len(phases), false)

	outputs := make([]int, 0)
	for _, phasePerm := range phasePermutations {
		outputs = append(outputs, runNetwork(instructions, phasePerm))
	}

	sort.Ints(outputs)

	// this is the answer to part 1
	fmt.Println(outputs[len(outputs)-1])

	// part 2
	phases2 := []int{5, 6, 7, 8, 9}
	phasePermutations2 := getPermutations(phases2, len(phases2), false)

	outputs2 := make([]int, 0)
	for _, phasePerm := range phasePermutations2 {
		outputs2 = append(outputs2, runFeedbackNetwork(instructions, phasePerm))
	}

	sort.Ints(outputs2)

	// this is the answer to part 2
	fmt.Println(outputs2[len(outputs2)-1])
}
