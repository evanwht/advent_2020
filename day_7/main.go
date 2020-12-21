package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nodeMap := NodeMap{nodes: make(map[string]*Node)}
	util.ForEachLine("day_7/input.txt", func(rule string) {
		parts := strings.Split(rule, " contain ")
		color := strings.Replace(parts[0], " bags", "", 1)
		node := nodeMap.getOrAdd(color)

		rules := strings.Split(parts[1], ", ")
		for _, r := range rules {
			if !strings.Contains(r, "no other bags") {
				count, err := strconv.Atoi(r[:1])
				util.Check(err)
				rColor := r[2:strings.Index(r, " bag")]
				n := nodeMap.getOrAdd(rColor)
				n.parents = append(n.parents, node)
				node.children[n] = count
			}
		}
	})
	fmt.Println(nodeMap.getContainedCount("shiny gold"))
	fmt.Println(nodeMap.getContainingCount("shiny gold"))
}

