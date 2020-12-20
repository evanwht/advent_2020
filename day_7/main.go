package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_7/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ruleSet := make(map[string]*Node)
	for scanner.Scan() {
		rule := scanner.Text()
		parts := strings.Split(rule, " contain ")
		color := parts[0][:strings.Index(parts[0], " bags")]
		rules := strings.Split(parts[1], ", ")
		node, ok := ruleSet[color]
		if !ok {
			node = &Node{
				color:    color,
				children: make(map[*Node]int),
			}
		}
		ruleSet[color] = node
		for _, r := range rules {
			if !strings.Contains(r, "no other bags") {
				count, err := strconv.Atoi(r[:1])
				util.Check(err)
				rColor := r[2:strings.Index(r, " bag")]
				if n, ok := ruleSet[rColor]; ok {
					n.parents = append(n.parents, node)
					node.children[n] = count
				} else {
					n = &Node{
						color:   rColor,
						parents: []*Node{node},
						children: make(map[*Node]int),
					}
					ruleSet[rColor] = n
					node.children[n] = count
				}
			}
		}
	}
	fmt.Println(getContainedCount(ruleSet, ruleSet["shiny gold"].parents))
	fmt.Println(getContainingCount(ruleSet, "shiny gold"))
}

func getContainedCount(ruleSet map[string]*Node, nodes []*Node) int {
	c := 0
	for _, n := range nodes {
		if r, ok := ruleSet[n.color]; ok && !r.counted {
			c++
			r.counted = true
		}
		c += getContainedCount(ruleSet, n.parents)
	}
	return c
}

func getContainingCount(ruleSet map[string]*Node, color string) int {
	c := 0
	if rule, ok := ruleSet[color]; ok {
		for n, i := range rule.children {
			c += i
			c += i * getContainingCount(ruleSet, n.color)
		}
	}
	return c
}

type Node struct {
	counted bool
	color string
	children map[*Node]int
	parents []*Node
}