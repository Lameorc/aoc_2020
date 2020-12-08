package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type programLine struct {
	instruction string
	val         int
}

type program map[int]programLine

func parseProgram(input []string) program {
	program := make(program)
	for i, l := range input {
		instrVal := strings.Split(l, " ")
		instr := instrVal[0]
		val, err := strconv.Atoi(instrVal[1])
		if err != nil {
			log.Fatal(err)
		}
		program[i] = programLine{instr, val}
	}

	return program
}

func runProgram(p program) (int, map[int]*programLine) {
	visited := make(map[int]*programLine)
	acc, i := 0, 0
	for {
		if i >= len(p) {
			break
		}
		p := p[i]
		visited[i] = &p
		if p.instruction == "nop" {
			i++
		} else if p.instruction == "jmp" {
			i += p.val
		} else if p.instruction == "acc" {
			acc += p.val
			i++
		}
		if visited[i] != nil {
			return acc, visited
		}
	}
	return acc, nil
}

func Solve(input []string) {
	baseProgram := parseProgram(input)
	part1, visited := runProgram(baseProgram)
	fmt.Printf("Part 1: %d\n", part1)

	swap := map[string]string{
		"jmp": "nop",
		"nop": "jmp",
		"acc": "acc", // no need to declare, will be skipped in this case
	}

	// just bruteforce it
	for i, line := range visited {
		newLine := programLine{line.instruction, line.val}
		if newLine.instruction == "acc" {
			continue
		} else {
			newLine.instruction = swap[newLine.instruction]
		}
		newProgram := make(program)

		// copy the original program and try the replacement
		for k, v := range baseProgram {
			if k == i {
				newProgram[k] = newLine
			} else {
				newProgram[k] = v
			}
		}
		part2, sortOfErr := runProgram(newProgram)
		if sortOfErr == nil {
			fmt.Printf("Part2 %d\n", part2)
			break
		}
	}
}
