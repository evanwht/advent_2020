package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	challenge1()
}

func challenge1() {
	valid := 0
	util.ForEachLine("input.txt", func(str string) {
		parts := strings.Split(str, ":")
		rule := parseRule(parts[0])
		pass := strings.TrimSpace(parts[1])
		if checkPassword2(pass, rule) {
			valid++
		}
	})
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
	util.Check(err)
	parts = strings.Split(parts[1], " ")
	m, err := strconv.Atoi(parts[0])
	util.Check(err)
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

func checkPassword2(password string, rule Rule) bool {
	first := rule.least <= len(password) && password[rule.least-1] == rule.letter
	second := rule.most <= len(password) && password[rule.most-1] == rule.letter
	return (first || second) && !(first && second)
}