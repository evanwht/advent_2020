package main

import (
	"advent_2020/util"
	"fmt"
)

func main() {
	treeMap := readMap()
	count := numTrees(treeMap, 1, 1)
	count *= numTrees(treeMap, 1, 3)
	count *= numTrees(treeMap, 1, 5)
	count *= numTrees(treeMap, 1, 7)
	count *= numTrees(treeMap, 2, 1)
	fmt.Println(count)
}

func numTrees(treeMap [][]byte, slopeX int, slopeY int) int {
	x, y, maxY, count := 0, 0, len(treeMap[0]), 0
	for x < len(treeMap) {
		if treeMap[x][y] == 35 {
			count++
		}
		x += slopeX
		y = (y + slopeY) % maxY
	}
	return count
}

func readMap() (treeMap [][]byte) {
	util.ForEachLine("input.txt", func(str string) {
		treeMap = append(treeMap, []byte(str))
	})
	return treeMap
}
