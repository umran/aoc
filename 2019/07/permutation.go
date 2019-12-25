package main

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
