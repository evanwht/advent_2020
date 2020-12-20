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
	t, counts := scanner.Text(), make(map[int32]struct{})
	for t != "" {
		// add this so we can see when readGroup didn't read
		counts[0] = struct{}{}

		for _, c := range t {
			counts[c] = struct{}{}
		}

		scanner.Scan()
		t = scanner.Text()
	}
	return len(counts) - 1
}

