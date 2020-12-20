package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day_5/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxSeatId := 0
	for scanner.Scan() {
		str := scanner.Text()
		if id := getSeatId(str); id > maxSeatId {
			maxSeatId = id
		}
	}
	fmt.Printf("max seat id: %d", maxSeatId)
}

func getSeatId(str string) int {
	row := getRow(str[:7])
	seat := getSeat(str[7:])
	return (row * 8) + seat
}

func getRow(rows string) int {
	row := 63
	for i, r := range rows {
		var change int
		if i < len(rows) - 1 {
			change = 32 >> i
		} else {
			change = 1
		}
		if r == 'F' {
			row -= change
		} else {
			row += change
		}
	}
	return row
}

func getSeat(seats string) int {
	seat := 4
	for i, s := range seats {
		var change int
		if i < len(seats) - 1 {
			change = 2 >> i
		} else {
			change = 1
		}
		if s == 'L' {
			seat -= change
		} else {
			seat += change
		}
	}
	return seat
}