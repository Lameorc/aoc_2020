package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type bus struct {
	index int
	mask  int
}

func newBuses(in string) ([]bus, error) {
	b := make([]bus, 0)
	for i, s := range strings.Split(in, ",") {
		if s == "x" {
			continue
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		b = append(b, bus{index: i, mask: val})
	}
	return b, nil
}

func part1(start int, buses []bus) int {
	leastWait := start
	leastWaitMask := 0
	for _, b := range buses {
		untilNext := b.untilNextFrom(start)
		if untilNext < leastWait {
			leastWait = untilNext
			leastWaitMask = b.mask
		}
	}
	return leastWait * leastWaitMask
}

func (b *bus) untilNextFrom(t int) int {
	busLeftAgo := t % b.mask
	untilNext := b.mask - busLeftAgo
	return untilNext
}

func part2(buses []bus) int {

	timestamp := 0
	// The timestamp interval must be between first and last inclusive, otherwise the first
	// leaves before the last one had a chance
	for ; ; timestamp++ {
		if timestamp%buses[0].mask != 0 {
			continue
		}
		busesConverging := 0
		for _, b := range buses {
			if (timestamp+b.index)%b.mask == 0 {
				busesConverging++
			} else {
				// no need now, we need all to match
				break
			}
		}
		if busesConverging == len(buses) {
			return timestamp
		}
	}
}

func Solve(input []string) {
	// first line is timestamp
	earliest, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}

	// second line the buses
	buses, err := newBuses(input[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(earliest, buses))
	fmt.Printf("Part 2: %d\n", part2(buses))
}
