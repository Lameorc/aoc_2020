package day12

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	c := shipCoordinates{facing: "E"}
	i := parseInstructions(input)

	result := part1(&c, i)
	if c.x != 17 {
		t.Errorf("expected x == 17, got %d", c.x)
	}
	if c.y != -8 {
		t.Errorf("expected x == 8, got %d", c.y)
	}
	if result != 25 {
		t.Errorf("expected result == 25, got %d", result)
	}
}

func TestDoInstr(t *testing.T) {
	instructions := map[instruction]shipCoordinates{
		{"L", 90}:  {0, 0, "N"},
		{"R", 90}:  {0, 0, "S"},
		{"R", 270}: {0, 0, "N"},
		{"L", 270}: {0, 0, "S"},
		{"R", 180}: {0, 0, "W"},
		{"L", 180}: {0, 0, "W"},
		{"F", 100}: {100, 0, "E"},
		{"N", 100}: {0, 100, "E"},
		{"E", 100}: {100, 0, "E"},
		{"S", 100}: {0, -100, "E"},
		{"W", 100}: {-100, 0, "E"},
	}
	for instr, expected := range instructions {
		ship := shipCoordinates{0, 0, "E"}
		ship.doInstr(instr)
		if ship != expected {
			t.Errorf("Expected %v, got %v after running %v", expected, ship, instr)
		}
	}
}

func TestForwardAfterTurn(t *testing.T) {
	s := shipCoordinates{0, 0, "E"}
	s.doInstr(instruction{"L", 90})
	s.doInstr(instruction{"F", 42})
	expected := shipCoordinates{0, 42, "N"}
	if s != expected {
		t.Errorf("Expected %v, got %v", expected, s)
	}

	s = shipCoordinates{0, 0, "E"}
	s.doInstr(instruction{"R", 90})
	s.doInstr(instruction{"F", 42})
	expected = shipCoordinates{0, -42, "S"}
	if s != expected {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}

func TestTurnMapping(t *testing.T) {
	for k, v := range rightTurnMap {
		if leftTurnMap[v] != k {
			t.Errorf("%v does not match it's coressponding left turn: %v", k, leftTurnMap[k])
		}
	}
}

func TestRotateWaypoint(t *testing.T) {
	w := waypoint{10, 30}
	w.rotateWaypoint("L", 180)
	if (w != waypoint{-10, -30}) {
		t.Fatalf("Rotating 180 left does not work as expected")
	}

	w = waypoint{10, 30}
	w.rotateWaypoint("R", 180)
	if (w != waypoint{-10, -30}) {
		t.Fatalf("Rotating 180 left does not work as expected")
	}

	w = waypoint{10, 30}
	w2 := waypoint{10, 30}
	w.rotateWaypoint("R", 90)
	w.rotateWaypoint("R", 90)
	w2.rotateWaypoint("R", 180)
	if w != w2 {
		t.Fatalf("Rotating right by 90 twice has different result than once by 180: %v <> %v", w, w2)
	}
}
