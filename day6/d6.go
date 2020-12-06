package day6

import (
	"fmt"
	"strings"
)

// I can't believe this is not in stdlib...
func intersect(a, b map[rune]bool) map[rune]bool {
	r := make(map[rune]bool)
	for k := range a {
		if b[k] {
			r[k] = true
		}
	}
	return r
}

func Solve(input []string) {
	part1 := 0
	part2 := 0

	answers := make(map[rune]bool)
	var allTrue map[rune]bool = nil
	for _, l := range input {
		if strings.TrimSpace(l) == "" {
			// going to next group
			part1 += len(answers)
			part2 += len(allTrue)
			answers = make(map[rune]bool)
			allTrue = nil
			continue
		}
		personsAnswers := make(map[rune]bool)
		for _, a := range l {
			answers[a] = true
			personsAnswers[a] = true
		}
		if allTrue == nil {
			allTrue = personsAnswers
		} else {
			allTrue = intersect(allTrue, personsAnswers)
		}

	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
