package day14

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type memory map[int]int

type programLine struct {
	instr string
	value string
}

type program []programLine

func newProgram(input []string) program {
	p := make(program, 0)
	for _, l := range input {
		instrValue := strings.Split(l, " = ")
		instr := instrValue[0]
		val := instrValue[1]
		if instr != "mask" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			// pad to 36
			val = fmt.Sprintf("%036b", intVal)
		}
		pl := programLine{instr: instr, value: val}
		p = append(p, pl)
	}

	return p
}

var memRe *regexp.Regexp = regexp.MustCompile(`mem\[(\d+)\]`)

func parseAddr(s string) int {
	memMatch := memRe.FindStringSubmatch(s)
	if len(memMatch) != 2 {
		log.Fatalf("failed memory address parsing")
	}
	memAddr, err := strconv.Atoi(memMatch[1])
	if err != nil {
		log.Fatal(err)
	}

	return memAddr
}

type writeInstruction func(m *memory, mask string, pl programLine)

func part1WriteInstr(m *memory, mask string, pl programLine) {
	// get the address and value
	memAddr := parseAddr(pl.instr)
	memVal := []rune(pl.value)
	// apply mask
	for i, v := range mask {
		if strings.ToUpper(string(v)) == "X" {
			continue
		}

		memVal[i] = v
	}

	// get the final value
	intVal, err := strconv.ParseInt(string(memVal), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	(*m)[memAddr] = int(intVal)
}

func part2WriteInstr(m *memory, mask string, pl programLine) {
	// get the address as binary
	memAddr := fmt.Sprintf("%036b", parseAddr(pl.instr))

	// the value can be taken as is, no mask application in this part
	memVal, err := strconv.ParseInt(string(pl.value), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	maskedAddr := []rune(memAddr)
	floatingAddresess := make(map[int]bool, 0)
	for i, r := range mask {
		if string(r) == "0" {
			continue
		}

		if strings.ToUpper(string(r)) == "X" {
			floatingAddresess[i] = true
		}
		maskedAddr[i] = r
	}

	nAddresses := int(math.Pow(2, float64(len(floatingAddresess))))
	addresses := make([]int, 0)

	addrMaskFmtStr := fmt.Sprintf("%%0%db", len(floatingAddresess))
	for i := 0; i < nAddresses; i++ {
		addrMask := fmt.Sprintf(addrMaskFmtStr, i)
		addr := make([]rune, len(maskedAddr))
		copy(addr, maskedAddr)
		for _, v := range addrMask {
			addr = []rune(strings.Replace(string(addr), "X", string(v), 1))
		}
		intVal, err := strconv.ParseInt(string(addr), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, int(intVal))
	}

	for _, a := range addresses {
		(*m)[a] = int(memVal)
	}
}

func (p *program) runProgram(m *memory, wi writeInstruction) {
	// program always should start with mask, so just take the first instruction
	currentMask := (*p)[0].value
	for _, i := range *p {
		if i.instr == "mask" {
			currentMask = i.value
		} else { // must be mem
			wi(m, currentMask, i)
		}
	}
}

func (m *memory) sumValues() int {
	sum := 0
	for _, v := range *m {
		sum += v
	}
	return sum
}

func part1(p program) int {
	mem := make(memory)
	p.runProgram(&mem, part1WriteInstr)

	return mem.sumValues()
}

func part2(p program) int {
	mem := make(memory)
	p.runProgram(&mem, part2WriteInstr)

	return mem.sumValues()

}

// Solve the day's problem
func Solve(input []string) {
	p := newProgram(input)
	log.Printf("Part 1: %d", part1(p))
	log.Printf("Part 2: %d", part2(p))

}
