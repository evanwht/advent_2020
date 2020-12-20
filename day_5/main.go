package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("day_5/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ids []int
	for scanner.Scan() {
		str := scanner.Text()
		ids = append(ids, getSeatId(str))
	}
	sort.Ints(ids)
	for i, id := range ids {
		if ids[i+1] == id+2 {
			fmt.Println(id+1)
		}
	}
}

func getSeatId(str string) int {
	row := getRow(str[:7])
	seat := getSeat(str[7:])
	return (row * 8) + seat
}

func getRow(rows string) int {
	row := 63.5
	for i, r := range rows {
		var change float64
		if i < len(rows) - 1 {
			shift := 32 >> i
			change = float64(shift)
		} else {
			change = 0.5
		}
		if r == 'F' {
			row -= change
		} else {
			row += change
		}
	}
	return int(row)
}

func getSeat(seats string) int {
	seat := 3.5
	for i, s := range seats {
		var change float64
		if i < len(seats) - 1 {
			shift := 2 >> i
			change = float64(shift)
		} else {
			change = 0.5
		}
		if s == 'L' {
			seat -= change
		} else {
			seat += change
		}
	}
	return int(seat)
}