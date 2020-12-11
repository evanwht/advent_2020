package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_4/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid, num := 0, 0
	// last one should have no fields filled in
	for pass := readPassport(scanner); !pass.isEmpty(); pass = readPassport(scanner) {
		 if pass.isValid() {
			 valid++
		 }
		 num++
	}
	fmt.Printf("%d of %d valid", valid, num)
}

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt int
	hgtUnit string
	hcl string
	ecl string
	pid string
	cid string // optional
}

func (p *Passport) isValid() bool {
	return p.byr >= 1920 && p.byr <= 2002 &&
		p.iyr >= 2010 && p.iyr <= 2020 &&
		p.eyr >= 2020 && p.eyr <= 2030 &&
		hgtValid(p.hgt, p.hgtUnit) &&
		hclValid(p.hcl) &&
		eclValid(p.ecl) &&
		pidValid(p.pid)
}

func pidValid(val string) bool {
	b, _ := regexp.Match("[0-9]+", []byte(val))
	return len(val) == 9 && b
}

func eclValid(val string) bool {
	return val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth"
}

func hclValid(val string) bool {
	b, _ := regexp.Match("#[0-9a-f]+", []byte(val))
	return len(val) == 7 && b
}

func hgtValid(val int, unit string) bool {
	if unit == "cm" {
		return val >= 150 && val <= 193
	} else if unit == "in" {
		return val >= 59 && val <= 76
	} else {
		return false
	}
}

func (p *Passport) isEmpty() bool {
	return p.byr == 0 &&
		p.iyr == 0 &&
		p.eyr == 0 &&
		p.hgt == 0 &&
		len(p.hgtUnit) == 0 &&
		len(p.hcl) == 0 &&
		len(p.ecl) == 0 &&
		len(p.pid) == 0 &&
		len(p.cid) == 0
}

func readPassport(scanner *bufio.Scanner) *Passport {
	scanner.Scan()
	t, passport := scanner.Text(), Passport{}
	for t != "" {
		parts := strings.Split(t, " ")
		for _, part := range parts {
			field := strings.Split(part, ":")
			switch field[0] {
			case "byr":
				// ignore error here because 0 will fail validation anyways
				passport.byr, _ = strconv.Atoi(field[1])
			case "iyr":
				passport.iyr, _ = strconv.Atoi(field[1])
			case "eyr":
				passport.eyr, _ = strconv.Atoi(field[1])
			case "hgt":
				passport.hgt, _ = strconv.Atoi(field[1][:len(field[1]) - 2])
				passport.hgtUnit = field[1][len(field[1]) - 2:]
			case "hcl":
				passport.hcl = field[1]
			case "ecl":
				passport.ecl = field[1]
			case "pid":
				passport.pid = field[1]
			case "cid":
				passport.cid = field[1]
			}
		}
		scanner.Scan()
		t = scanner.Text()
	}
	return &passport
}