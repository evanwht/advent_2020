package main

import (
	"advent_2020/util"
	"fmt"
	"sort"
)

func main() {
	var ids []int
	util.ForEachLine("day_5/input.txt", func(str string) {
		ids = append(ids, getSeatId(str))
	})
	sort.Ints(ids)
	for i, id := range ids {
		if i < len(ids)-1 && ids[i+1] == id+2 {
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
	row := 0
	for i, r := range rows {
		if r == 'B' {
			row |= 1 << (len(rows) - (i+1))
		}
	}
	return row
}

func getSeat(seats string) int {
	seat := 0
	for i, s := range seats {
		if s == 'R' {
			seat |= 1 << (len(seats) - (i+1))
		}
	}
	return seat
}