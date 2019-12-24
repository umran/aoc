package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	pwMin = 372304
	pwMax = 847060
)

func hasSixDigits(pw int) bool {
	pwString := strconv.FormatInt(int64(pw), 10)

	slices := strings.Split(pwString, "")
	return len(slices) == 6
}

func hasDouble(pw int) bool {
	pwString := strconv.FormatInt(int64(pw), 10)

	slices := strings.Split(pwString, "")
	for i := 0; i < len(slices)-1; i++ {
		if slices[i] == slices[i+1] {
			return true
		}
	}

	return false
}

func hasIsolatedDouble(pw int) bool {
	pwString := strconv.FormatInt(int64(pw), 10)

	slices := strings.Split(pwString, "")

	seenList := make(map[string]int)
	lastSeen := ""
	for i := 0; i < len(slices); i++ {
		slice := slices[i]
		if slice != lastSeen {
			if seenList[lastSeen] == 2 {
				return true
			}

			seenList[slice] = 1

			// reset previous count
			seenList[lastSeen] = 0
		} else {
			seenList[slice] = seenList[slice] + 1
			if i == len(slices)-1 {
				if count := seenList[slice]; count == 2 {
					return true
				}
			}
		}

		lastSeen = slice
	}

	return false
}

func isIncreasing(pw int) bool {
	pwString := strconv.FormatInt(int64(pw), 10)

	slices := strings.Split(pwString, "")
	for i := 0; i < len(slices)-1; i++ {
		a, _ := strconv.ParseInt(slices[i], 10, 64)
		b, _ := strconv.ParseInt(slices[i+1], 10, 64)
		if a > b {
			return false
		}
	}

	return true
}

func isPwCandidate(pw int) bool {
	return hasSixDigits(pw) && hasDouble(pw) && isIncreasing(pw)
}

func isRevisedPwCandidate(pw int) bool {
	return hasSixDigits(pw) && hasIsolatedDouble(pw) && isIncreasing(pw)
}

func main() {
	pwCandidates := make([]int, 0)
	for i := 0; i < pwMax-pwMin+1; i++ {
		pw := pwMin + i
		if isPwCandidate(pw) {
			pwCandidates = append(pwCandidates, pw)
		}
	}

	// this is the answer to part 1
	fmt.Println(len(pwCandidates))

	revisedPwCandidates := make([]int, 0)
	for i := 0; i < pwMax-pwMin+1; i++ {
		pw := pwMin + i
		if isRevisedPwCandidate(pw) {
			revisedPwCandidates = append(revisedPwCandidates, pw)
		}
	}

	// this is the answer to part 2
	fmt.Println(len(revisedPwCandidates))
}
