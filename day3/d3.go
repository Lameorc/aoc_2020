package day3

import "log"

type forest struct {
	rows   []string
	width  int
	length int
}

// creates the forest struct from puzzle input
func createForest(input []string) forest {
	forestLen := len(input)
	forestRows := make([]string, forestLen)
	for i, row := range input {
		forestRows[i] = row
	}
	f := forest{forestRows, len(forestRows[0]), forestLen}

	return f
}

type myPosition struct {
	x int
	y int
}

func (m *myPosition) movePos(s *slope, f *forest) {
	newX, newY := m.x+s.vx, m.y+s.vy
	// at the bottom
	if newY >= f.length {
		return
	}
	m.y = newY

	if newX >= f.width {
		newX = newX - f.width
	}
	m.x = newX
}

func (m *myPosition) checkTreeColission(f *forest) int {
	if f.rows[m.y][m.x] == '#' {
		return 1
	}
	return 0
}

type slope struct {
	vx int
	vy int
}

// ...
func Solve(input []string) int {
	total := 1
	f := createForest(input)
	// part 1
	// slopes := [1]slope{slope{3, 1}}

	// part 2
	slopes := [5]slope{
		slope{1, 1},
		slope{3, 1},
		slope{5, 1},
		slope{7, 1},
		slope{1, 2},
	}

	for _, s := range slopes {
		collissions := 0
		p := myPosition{0, 0}
		for {
			prevY := p.y
			p.movePos(&s, &f)

			// did not move down, so we're at the bottom
			if prevY == p.y {
				// update minimum if needed
				log.Printf("Colissions at slope <%d:%d>: %d", s.vx, s.vy, collissions)
				total *= collissions
				break
			}
			collissions += p.checkTreeColission(&f)
		}

	}

	return total
}
