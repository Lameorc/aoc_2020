package day9

import (
	"log"
	"math"
	"strconv"
)

type queue struct {
	elements []int
	maxLen   int
}

func (q *queue) push(e int) {
	q.elements = append(q.elements, e)
	if len(q.elements) > q.maxLen {
		q.elements = q.elements[1:]
	}
}

// determines whether the sum arg is a sum of two numbers from in
func isSumFromArr(sum int, in []int) bool {
	seen := make(map[int]int)
	for _, i := range in {
		if seen[sum-i]+i == sum {
			return true
		}
		seen[i] = i
	}

	return false
}

func part1(input []int) int {
	preambleSize := 25
	q := queue{elements: make([]int, 0), maxLen: preambleSize}
	i := 0
	for ; i < preambleSize; i++ {
		q.push(input[i])
	}

	for ; i < len(input); i++ {
		n := input[i]
		if !isSumFromArr(n, q.elements) {
			return n
		}
		q.push(n)
	}
	return 0
}

func part2(toSum int, in []int) int {
	for i := 0; in[i] != toSum; i++ {
		walkingSum := in[i]
		min, max := float64(in[i]), float64(in[i])
		for j := i + 1; walkingSum < toSum || j < len(in); j++ {
			newSum := walkingSum + in[j]
			min = math.Min(min, float64(in[j]))
			max = math.Max(max, float64(in[j]))
			if newSum == toSum {
				return int(min) + int(max)
			}
			walkingSum = newSum
		}

	}
	return 0
}

func Solve(input []string) {
	asInts := make([]int, 0)
	for _, l := range input {
		i, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		asInts = append(asInts, i)
	}
	p1 := part1(asInts)
	log.Printf("Part 1: %d", p1)
	p2 := part2(p1, asInts)
	log.Printf("Part 2: %d", p2)

}
