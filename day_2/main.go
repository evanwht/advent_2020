package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	challenge1()
}

func challenge1() {
	file, err := os.Open("day_2.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid := 0
	for scanner.Scan() {
		str := scanner.Text()
		parts := strings.Split(str, ":")
		rule := parseRule(parts[0])
		pass := strings.TrimSpace(parts[1])
		if checkPassword(pass, rule) {
			valid++
		}
	}
	fmt.Println(valid)
}

// Rule validation rule for a password
type Rule struct {
	least  int
	most   int
	letter byte
}

func parseRule(rule string) Rule {
	parts := strings.Split(rule, "-")
	l, err := strconv.Atoi(parts[0])
	check(err)
	parts = strings.Split(parts[1], " ")
	m, err := strconv.Atoi(parts[0])
	check(err)
	bytes := []byte(parts[1])
	return Rule{l, m, bytes[0]}
}

func checkPassword(password string, rule Rule) bool {
	count := 0
	for i := 0; i < len(password); i++ {
		if password[i] == rule.letter {
			count++
			if count > rule.most {
				return false
			}
		}
	}
	return count >= rule.least
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
