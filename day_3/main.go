package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	treeMap := readMap()
	x, y, maxY, count := 0, 0, len(treeMap[0]), 0
	for x < len(treeMap) {
		if treeMap[x][y] == 35 {
			count++
		}
		x++
		y = (y + 3) % maxY
	}
	fmt.Println(count)
}

func readMap() (treeMap [][]byte) {
	file, err := os.Open("input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		treeMap = append(treeMap, []byte(scanner.Text()))
	}
	return treeMap
}
