package day4

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
		"",
	}
	p := createPassports(input)
	if len(p) != 4 {
		t.Errorf("Number of passports is incorrect, expected 4, got %d", len(p))
	}
	valid := 0
	for i, pass := range p {
		if pass.isValidPart1() {
			valid++
			if i != 0 && i != 2 {
				t.Errorf("Did not expect passpot no.%d to be valid", i)
			}
		}
	}
	if valid != 2 {
		t.Errorf("Incorrect number of valid passports, expected 2, got %d", valid)
	}
}
