package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var ruleRe = regexp.MustCompile("(.+) bags contain ((?:\\d[^\\.,]*[,\\.]\\W?)*)")
var emptyBagRe = regexp.MustCompile("(.+) bags contain no other bags.")
var bagCountRe = regexp.MustCompile(`(\d) (.*) bags?\.?`)

func part1(input []string) int {

	possibleBags := map[string]bool{
		"shiny gold": true,
	}
	lenPos := 0

	// just iterate again and again like a monkey
	for lenPos != len(possibleBags) {
		lenPos = len(possibleBags)
		for _, l := range input {
			ms := ruleRe.FindStringSubmatch(l)
			if len(ms) != 3 {
				continue
			}
			bagName := ms[1]
			bagContents := ms[2]

			// end early here
			if possibleBags[bagName] {
				continue
			}

			if strings.Contains(bagContents, "shiny gold") {
				possibleBags[bagName] = true
			}
			for k := range possibleBags {
				if strings.Contains(bagContents, k) {
					possibleBags[bagName] = true
				}
			}
		}
	}

	// -1 since we were counting even our requested bag
	return lenPos - 1
}

func recurseRules(input []string, toCheck string) int {
	bagsToCheck := make(map[string]int)
	sum := 0
	for _, l := range input {
		bagIsEmptyMatch := emptyBagRe.FindStringSubmatch(l)
		if len(bagIsEmptyMatch) == 2 {
			if strings.TrimSpace(bagIsEmptyMatch[1]) == toCheck {
				return 0 // end recursion
			}
		}

		ms := ruleRe.FindStringSubmatch(l)
		if len(ms) != 3 {
			continue
		}
		bagName := strings.TrimSpace(ms[1])
		if bagName == toCheck {
			rules := strings.Split(ms[2], ",")
			for _, r := range rules {
				m := bagCountRe.FindStringSubmatch(r)
				if len(m) != 3 {
					log.Fatalf("Incorrectly formatted rule part: %s", r)
				}
				v, _ := strconv.Atoi(m[1])
				containedBag := strings.TrimSpace(m[2])
				bagsToCheck[containedBag] = v
			}
		}

	}
	for bag, count := range bagsToCheck {
		sum += count + (count * (recurseRules(input, bag)))
	}
	return sum
}

func part2(input []string) int {
	bagsNeeded := recurseRules(input, "shiny gold")
	return bagsNeeded
}

// Solve ...
func Solve(input []string) {
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))

}
