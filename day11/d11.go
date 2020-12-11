package day11

import (
	"fmt"
	"log"
)

// convert to slice of slices for easier indexing
type seatingArea [][]rune

func newSeatingArea(input []string) (s seatingArea) {
	s = make(seatingArea, 0)
	for _, r := range input {
		s = append(s, []rune(r))
	}
	return
}

func (s seatingArea) occupiedSeats() int {
	o := 0
	for _, r := range s {
		for _, c := range r {
			if c == '#' {
				o++
			}
		}
	}
	return o
}

func (s seatingArea) isEqual(other seatingArea) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if other[i][j] != s[i][j] {
				return false
			}
		}
	}
	return true
}

func (s seatingArea) print() {
	fmt.Printf("xSize: %d; ySize: %d\n", len(s), len(s[0]))
	for _, r := range s {
		fmt.Printf("\t%s\n", string(r))
	}
}

func tick(s seatingArea) seatingArea {
	// make a copy modifying as needed
	maxX := len(s)
	newS := make(seatingArea, maxX)
	for i, r := range s {
		maxY := len(r)
		row := make([]rune, maxY)
		for j := range r {
			row[j] = getNewTileValue(&s, i, j, maxX, maxY)
		}
		newS[i] = row
	}

	return newS
}

func getNewTileValue(s *seatingArea, x, y, maxX, maxY int) rune {
	initial := (*s)[x][y]
	if initial == '.' {
		return '.'
	}

	adjacentOccupied := 0

	// assummes symetrical
	visionRangeX, visionRangeY := 1, 1 // part1
	//visionRange := maxX

	for k := -visionRangeX; k <= visionRangeX; k++ {
		newX := k + x
		// bounds
		if newX < 0 || newX >= maxX {
			continue
		}

		for l := -visionRangeY; l <= visionRangeY; l++ {
			newY := l + y
			// don't check self
			if x == newX && y == newY {
				continue
			}
			// bounds
			if newY < 0 || newY >= maxY {
				continue
			}

			// finally check the tile
			if (*s)[newX][newY] == '#' {
				adjacentOccupied++
			}
		}
	}

	var n rune
	if initial == '#' && adjacentOccupied >= 4 {
		n = 'L'
	} else if initial == 'L' && adjacentOccupied == 0 {
		n = '#'
	} else {
		n = initial
	}

	return n
}

// Solve the day
func Solve(input []string) {
	s := newSeatingArea(input)
	ticks := 0
	for ; ; ticks++ {
		nextS := tick(s)
		if nextS.isEqual(s) {
			break
		} else {
			s = nextS
		}
	}

	log.Printf("Part 1: %d", s.occupiedSeats())
}
