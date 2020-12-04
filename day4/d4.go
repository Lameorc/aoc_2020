package day4

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func createPassports(input []string) []passport {
	var ps []passport

	p := passport{}
	for _, l := range input {
		if strings.TrimSpace(l) == "" {
			// have all the passport data available
			ps = append(ps, p)
			p = passport{}
			continue
		}
		pVal := reflect.ValueOf(&p).Elem()
		for _, v := range strings.Split(l, " ") {
			kv := strings.Split(v, ":")
			// need to convert first letter to uppercase due to golang export rules
			keyRunes := []rune(kv[0])
			keyRunes[0] = unicode.ToUpper(keyRunes[0])
			kv[0] = string(keyRunes)

			pVal.FieldByName(kv[0]).SetString(kv[1])
		}
	}

	return ps
}

func (p *passport) isValidPart1() bool {
	s := reflect.ValueOf(p).Elem()
	typeOfS := s.Type()
	allPresent := true
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if typeOfS.Field(i).Name == "Cid" {
			continue
		}

		allPresent = allPresent && (f.Interface() != "")

	}
	return allPresent

}

func checkYear(v string, min int, max int) bool {
	asInt, _ := strconv.Atoi(v)
	return len(v) == 4 && asInt >= min && asInt <= max
}

func checkHeight(v string) bool {
	if len(v) <= 2 {
		return false
	}
	unit := v[len(v)-2:]
	val, err := strconv.Atoi(v[:len(v)-2])
	if err != nil {
		return false
	}

	if unit == "cm" {
		return val >= 150 && val <= 193
	}

	if unit == "in" {
		return val >= 59 && val <= 76
	}

	// invalid unit
	return false
}

func (p *passport) isValidPart2() bool {
	log.Printf("Checking %v", p)
	if !checkYear(p.Byr, 1920, 2002) {
		log.Println("Not valid due to byr")
		return false
	}

	if !checkYear(p.Iyr, 2010, 2020) {
		log.Println("Not valid due to iyr")
		return false
	}
	if !checkYear(p.Eyr, 2020, 2030) {
		log.Println("Not valid due to eyr")
		return false
	}

	if !checkHeight(p.Hgt) {
		return false
	}

	asRunes := []rune(p.Hcl)
	_, err := strconv.ParseInt(string(asRunes[1:]), 16, 32)
	if string(asRunes[0]) != "#" || err != nil || len(p.Hcl) != 7 {
		log.Println("Not valid due to hcl")
		return false
	}

	validEcl := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}
	if !validEcl[p.Ecl] {
		log.Println("Not valid due to ecl")
		return false
	}
	_, err = strconv.Atoi(p.Pid)
	if len(p.Pid) != 9 || err != nil {
		log.Println("Not valid due to pid")
		return false
	}

	log.Println("Valid!")
	return true
}

func part1(ps []passport) {
	valid1 := 0
	valid2 := 0
	for _, p := range ps {
		if p.isValidPart1() {
			valid1++
		}
		if p.isValidPart2() {
			valid2++
		}
	}

	fmt.Printf("Part 1: %d", valid1)
	fmt.Printf("Part 2: %d", valid2)
}

func Solve(input []string) {
	p := createPassports(input)
	part1(p)

}
