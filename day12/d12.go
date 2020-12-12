package day12

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

type direction struct {
	dx, dy int
}

var directions = map[string]direction{
	"N": {0, 1},
	"S": {0, -1},
	"E": {1, 0},
	"W": {-1, 0},
}

var directionToName = map[direction]string{
	directions["N"]: "N",
	directions["S"]: "S",
	directions["E"]: "E",
	directions["W"]: "W",
}

var rightTurnMap = map[direction]direction{
	directions["N"]: directions["E"],
	directions["E"]: directions["S"],
	directions["S"]: directions["W"],
	directions["W"]: directions["N"],
}

// inverse to right turns
var leftTurnMap = map[direction]direction{
	directions["N"]: directions["W"],
	directions["W"]: directions["S"],
	directions["S"]: directions["E"],
	directions["E"]: directions["N"],
}

type shipCoordinates struct {
	x, y   int
	facing string
}

func (s *shipCoordinates) doInstr(i instruction) {
	log.Printf("running %v", i)
	switch a := i.action; a {
	case "N", "S", "E", "W":
		s.x += directions[a].dx * i.value
		s.y += directions[a].dy * i.value
	case "F":
		s.x += directions[s.facing].dx * i.value
		s.y += directions[s.facing].dy * i.value
	case "L":
		switch current := directions[s.facing]; i.value {
		case 180:
			s.facing = directionToName[direction{-current.dx, -current.dy}]
		case 90:
			s.facing = directionToName[leftTurnMap[direction{current.dx, current.dy}]]
		// same as 90 opposite direction
		case 270:
			s.facing = directionToName[rightTurnMap[direction{current.dx, current.dy}]]
		default:
			log.Fatalf("Unknown turn value %d", i.value)
		}
	case "R":
		switch current := directions[s.facing]; i.value {
		case 180:
			s.facing = directionToName[direction{-current.dx, -current.dy}]
		case 90:
			s.facing = directionToName[rightTurnMap[direction{current.dx, current.dy}]]
		// same as 90 opposite direction
		case 270:
			s.facing = directionToName[leftTurnMap[direction{current.dx, current.dy}]]
		default:
			log.Fatalf("Unknown turn value %d", i.value)
		}
	default:
		log.Fatalf("Unknown instruction %s", a)
	}
}

func (s *shipCoordinates) moveToWaypoint(w *waypoint) {
	s.x += w.x
	s.y += w.y
}

type waypoint struct {
	x, y int
}

func (w *waypoint) moveWaypoint(i instruction) {
	log.Printf("Moving waypoint by %v", i)
	w.x += directions[i.action].dx * i.value
	w.y += directions[i.action].dy * i.value
}
func (w *waypoint) rotateWaypoint(direction string, angle int) {
	if direction == "L" {
		switch angle {
		case 90:
			w.x, w.y = -w.y, w.x
		case 180:
			w.x, w.y = -w.x, -w.y
		case 270:
			w.x, w.y = w.y, -w.x
		}
	}
	if direction == "R" {
		switch angle {
		case 90:
			w.x, w.y = w.y, -w.x
		case 180:
			w.x, w.y = -w.x, -w.y
		case 270:
			w.x, w.y = -w.y, w.x
		}
	}
}

type instruction struct {
	action string
	value  int
}

var reInstr = regexp.MustCompile(`(\w)(\d+)`)

func parseInstructions(input []string) []instruction {
	i := make([]instruction, 0)
	for _, l := range input {
		actionvalue := reInstr.FindStringSubmatch(l)
		if len(actionvalue) != 3 {
			log.Fatalf("failed to parse %s, result was %v", l, actionvalue)
		}
		action := actionvalue[1]
		value, err := strconv.Atoi(actionvalue[2])
		if err != nil {
			log.Fatal(err)
		}
		instr := instruction{action, value}
		i = append(i, instr)
	}

	return i
}

func part1(c *shipCoordinates, i []instruction) int {
	for _, instr := range i {
		log.Printf("Prev cords %s:[%d, %d]", c.facing, c.x, c.y)
		c.doInstr(instr)
		log.Printf("New cords %s:[%d, %d]", c.facing, c.x, c.y)
	}
	return int(math.Abs(float64(c.x))) + int(math.Abs(float64(c.y)))
}

func part2(c *shipCoordinates, i []instruction) int {
	w := waypoint{10, 1}
	for _, instr := range i {
		log.Printf("Before : Ship[%d, %d], waypoint[%d, %d]", c.x, c.y, w.x, w.y)
		switch instr.action {
		case "N", "E", "S", "W":
			w.moveWaypoint(instr)
		case "L", "R":
			w.rotateWaypoint(instr.action, instr.value)
		case "F":
			log.Printf("moving ship to waypoint%v %d times", w, instr.value)
			for i := 0; i < instr.value; i++ {
				c.moveToWaypoint(&w)
			}
		}
		log.Printf("After: Ship[%d, %d], waypoint[%d, %d]", c.x, c.y, w.x, w.y)
	}
	return int(math.Abs(float64(c.x))) + int(math.Abs(float64(c.y)))

}

// Solve the day
func Solve(input []string) {
	c := shipCoordinates{facing: "E"}
	i := parseInstructions(input)

	log.Printf("Part1: %d", part1(&c, i))

	c = shipCoordinates{facing: "E"}
	log.Printf("Part2: %d", part2(&c, i))

}
