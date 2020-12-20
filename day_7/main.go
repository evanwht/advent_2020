package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
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
		parts := strings.Split(rule, "contain")
		color := parts[0][:strings.Index(parts[0], " bags")]
		rules := strings.Split(parts[1], ",")
		node, ok := ruleSet[color]
		if !ok {
			node = &Node{
				color:    color,
			}
		}
		ruleSet[color] = node
		for _, r := range rules {
			if !strings.Contains(r, "no other bags") {
				rColor := r[3:strings.Index(r, " bag")]
				if n, ok := ruleSet[rColor]; ok {
					n.parents = append(n.parents, node)
					node.children = append(node.children, n)
				} else {
					n = &Node{
						color:   rColor,
						parents: []*Node{node},
					}
					ruleSet[rColor] = n
					node.children = append(node.children, n)
				}
			}
		}
	}
	fmt.Println(getCount(ruleSet, ruleSet["shiny gold"].parents))
}

func getCount(ruleSet map[string]*Node, nodes []*Node) int {
	c := 0
	for _, n := range nodes {
		if r, ok := ruleSet[n.color]; ok && !r.counted {
			c++
			r.counted = true
		}
		c += getCount(ruleSet, n.parents)
	}
	return c
}

type Node struct {
	counted bool
	color string
	children []*Node
	parents []*Node
}