package day5

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type boardingPass struct {
	row    int
	column int
	id     int
}

func passFromInput(in string) *boardingPass {

	// F == 0, B == 1
	asBinaryStr := strings.ReplaceAll(in, "F", "0")
	asBinaryStr = strings.ReplaceAll(asBinaryStr, "B", "1")
	// L == 0, R == 1
	asBinaryStr = strings.ReplaceAll(asBinaryStr, "L", "0")
	asBinaryStr = strings.ReplaceAll(asBinaryStr, "R", "1")

	row, err := strconv.ParseInt(asBinaryStr[0:7], 2, 16)
	if err != nil {
		log.Fatal(err)
	}

	column, err := strconv.ParseInt(asBinaryStr[7:], 2, 16)
	if err != nil {
		log.Fatal(err)
	}

	id := (row * 8) + column

	return &boardingPass{row: int(row), column: int(column), id: int(id)}
}

// Solve ...
func Solve(input []string) {
	presentPasses := make(map[int]bool)
	var max float64
	var min float64

	for _, l := range input {
		pass := passFromInput(l)
		presentPasses[pass.id] = true
		max = math.Max(max, float64(pass.id))
		if min == 0 { // first pass
			min = max
		} else {
			min = math.Min(min, float64(pass.id))
		}
	}
	fmt.Printf("Part 1: %d\n", int(max))

	// just linearly...
	for i := int(min); i <= int(max); i++ {
		if !presentPasses[i] {
			fmt.Printf("Part 2: %d\n", i)
			return
		}
	}

}
