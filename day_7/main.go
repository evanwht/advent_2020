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
	nodeMap := NodeMap{nodes: make(map[string]*Node)}
	for scanner.Scan() {
		rule := scanner.Text()
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
	}
	fmt.Println(nodeMap.getContainedCount("shiny gold"))
	fmt.Println(nodeMap.getContainingCount("shiny gold"))
}

