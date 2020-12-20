package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day_6/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int
	for c := readGroup(scanner); c > -1; c = readGroup(scanner) {
		count += c
	}
	fmt.Println(count)
}

func readGroup(scanner *bufio.Scanner) int {
	scanner.Scan()
	t, counts, group, count := scanner.Text(), make(map[int32]int), 0, 0
	for t != "" {

		// add this so we can see when readGroup didn't read
		count=1
		counts[0] = 0

		for _, c := range t {
			counts[c]++
		}

		group++
		scanner.Scan()
		t = scanner.Text()
	}
	for _, c := range counts {
		if c == group {
			count++
		}
	}
	return count - 1
}

