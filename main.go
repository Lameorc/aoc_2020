package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/lameorc/aoc_2020/day10"
)

func readInput(day string) []string {
	filePath := fmt.Sprintf("./%s/input_test.txt", day)
	// Should fit in memory easily
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

// Splits a given input into two dimensional slice, the first dimension is a group and the
// second are its elements. Expects newline after the last one
func toGroups(input []string) [][]string {
	groups := make([][]string, 0)
	currentGroup := make([]string, 0)
	for _, l := range input {
		if strings.TrimSpace(l) == "" {
			// going to next group
			groups = append(groups, currentGroup)
			currentGroup = make([]string, 0)
			continue
		}
		currentGroup = append(currentGroup, l)
	}
	return groups
}

func main() {
	// day1
	// in := readInput("day1")
	// fmt.Println(day1.Solve(in))

	// day2
	// in := readInput("day2")
	// fmt.Println(day2.Solve(in))

	// day3
	// in := readInput("day3")
	// fmt.Println(day3.Solve(in))

	// day4
	// in := readInput("day4")
	// day4.Solve(in)

	// day5
	// in := readInput("day5")
	// day5.Solve(in)

	// day 6
	// in := readInput("day6")
	// day6.Solve(in)

	// day 7
	// in := readInput("day7")
	// day7.Solve(in)

	// day 8
	// in := readInput("day8")
	// day8.Solve(in)

	// day 9
	// in := readInput("day9")
	// day9.Solve(in)

	// day 10
	in := readInput("day10")
	day10.Solve(in)
}
