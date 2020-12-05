package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/lameorc/aoc_2020/day5"
)

func readInput(day string) []string {
	filePath := fmt.Sprintf("./%s/input.txt", day)
	// Should fit in memory easily
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
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
	in := readInput("day5")
	day5.Solve(in)
}
