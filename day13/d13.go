package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type bus struct {
	mask int
}

func newBuses(in string) ([]bus, error) {
	b := make([]bus, 0)
	for _, s := range strings.Split(in, ",") {
		if s == "x" {
			continue
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		b = append(b, bus{val})
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
}
