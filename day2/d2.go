package day2

import (
	"log"
	"strconv"
	"strings"
)

type policy struct {
	// first index -1 in part 2
	min int
	// second index -1 in part 2
	max    int
	letter rune
}

// the input is expected to be in format of <min>-<max><space><letter>
func makePolicyFromInput(input string) (*policy, error) {
	// parts are range and letter
	parts := strings.Split(input, " ")
	minMaxRange := strings.Split(parts[0], "-")

	min, err := strconv.ParseInt(minMaxRange[0], 0, 32)
	if err != nil {
		return nil, err
	}
	max, err := strconv.ParseInt(minMaxRange[1], 0, 32)
	if err != nil {
		return nil, err
	}

	p := policy{int(min), int(max), rune(strings.TrimSpace(parts[1])[0])}

	return &p, nil
}

func Solve(lines []string) int {
	solution := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")

		policy, err := makePolicyFromInput(parts[0])
		if err != nil {
			log.Fatal(err)
			continue
		}

		password := strings.TrimSpace(parts[1])

		// part 1
		// nChars := 0
		// for _, c := range password {
		// 	if c == policy.letter {
		// 		nChars++
		// 	}
		// 	if nChars > policy.max {
		// 		// not valid, we already have more than max
		// 		break
		// 	}
		// }

		// if nChars < policy.min || nChars > policy.max {
		// 	// not valid not enough
		// 	continue
		// }
		// // getting here means it's valid
		// solution++

		// part 2
		passAsRunes := []rune(password)
		if (passAsRunes[policy.min-1] == policy.letter) != (passAsRunes[policy.max-1] == policy.letter) {
			solution++
		}

	}

	return solution
}
