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
	newS := make(seatingArea, len(s))
	for i, r := range s {
		row := make([]rune, len(r))
		for j := range r {
			row[j] = getNewTileValue(&s, i, j)
		}
		newS[i] = row
	}

	return newS
}

type direction [2]int

var directions = map[string]direction{
	"N":  {0, -1},
	"S":  {0, 1},
	"E":  {1, 0},
	"W":  {-1, 0},
	"NE": {1, -1},
	"NW": {-1, -1},
	"SE": {1, 1},
	"SW": {-1, 1},
}

func (s *seatingArea) directionOccupied(d direction, x, y, sight int) bool {
	maxX := len(*s)
	maxY := len((*s)[0])
	for i := 1; i <= sight; i++ {
		newX := x + (d[0] * i)
		if newX < 0 || newX >= maxX {
			return false
		}
		newY := y + (d[1] * i)
		// don't check self
		if x == newX && y == newY {
			return false
		}
		// bounds
		if newY < 0 || newY >= maxY {
			return false
		}

		// finally check the tile
		switch (*s)[newX][newY] {
		case '#':
			return true
		case 'L':
			return false
		default:
			continue
		}
	}
	return false
}

func (s *seatingArea) occupiedAdjacent(x, y int) int {
	adjacentOccupied := 0

	for _, modifiers := range directions {
		// if s.directionOccupied(modifiers, x, y, 1) { // ~part_1
		if s.directionOccupied(modifiers, x, y, 100) {
			adjacentOccupied++
		}
	}
	return adjacentOccupied
}

func getNewTileValue(s *seatingArea, x, y int) rune {
	initial := (*s)[x][y]
	if initial == '.' {
		return '.'
	}
	adjacentOccupied := s.occupiedAdjacent(x, y)
	var n rune
	// if initial == '#' && adjacentOccupied >= 4 { // ~part_1
	if initial == '#' && adjacentOccupied >= 5 {
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

	log.Printf("Part 2: %d", s.occupiedSeats())
}
