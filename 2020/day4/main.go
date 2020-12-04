package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	passports := parse(input)

	part1, part2 := 0, 0

	//byr (Birth Year)
	//iyr (Issue Year)
	//eyr (Expiration Year)
	//hgt (Height)
	//hcl (Hair Color)
	//ecl (Eye Color)
	//pid (Passport ID)
	//cid (Country ID)

	for _, passport := range passports {
		count := 0
		count2 := 0
		for _, expected := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			value, exist := passport[expected]
			if !exist {
				continue
			}
			count++
			switch expected {
			case "byr":
				nb, err := strconv.Atoi(value)
				if err == nil && nb >= 1920 && nb <= 2020 {
					count2++
				}
			case "iyr":
				nb, err := strconv.Atoi(value)
				if err == nil && nb >= 2010 && nb <= 2020 {
					count2++
				}
			case "eyr":
				nb, err := strconv.Atoi(value)
				if err == nil && nb >= 2020 && nb <= 2030 {
					count2++
				}
			case "hgt":
				hasCm := strings.Contains(value, "cm")
				hasIn := strings.Contains(value, "in")
				if hasCm || hasIn {
					nb, err := strconv.Atoi(value[:len(value)-2])
					if err != nil {
						continue
					}
					if hasCm && nb >= 150 && nb <= 193 || hasIn && nb >= 59 && nb <= 76 {
						count2++
					}
				}
			case "hcl":
				if regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(value) {
					count2++
				}
			case "ecl":
				switch value {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					count2++
				}
			case "pid":
				if regexp.MustCompile(`^[0-9]{9}$`).MatchString(value) {
					count2++
				}
			}
		}
		if count == 7 {
			part1++
		}
		if count2 == 7 {
			part2++
		}
	}

	return part1, part2
}

func parse(s string) []map[string]string {
	passportLines := strings.Split(s, "\n\n")
	passports := []map[string]string(nil)
	for i, pl := range passportLines {
		passports = append(passports, make(map[string]string))
		for _, line := range strings.Split(pl, "\n") {
			for _, attr := range strings.Split(line, " ") {
				tmp := strings.Split(attr, ":")
				passports[i][tmp[0]] = tmp[1]
			}
		}
	}
	return passports
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
