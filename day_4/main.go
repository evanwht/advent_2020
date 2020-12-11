package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
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
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string // optional
}

func (p *Passport) isValid() bool {
	return len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0
}

func (p *Passport) isEmpty() bool {
	return len(p.byr) == 0 &&
		len(p.iyr) == 0 &&
		len(p.eyr) == 0 &&
		len(p.hgt) == 0 &&
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
				passport.byr = field[1]
			case "iyr":
				passport.iyr = field[1]
			case "eyr":
				passport.eyr = field[1]
			case "hgt":
				passport.hgt = field[1]
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