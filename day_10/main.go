package main

import (
	"advent_2020/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	challenge1()
}

func challenge1() {
	var jolts []int
	util.ForEachLine("day_10/input.txt", func(str string) {
		n, _ := strconv.Atoi(str)
		jolts = append(jolts, n)
	})
	sort.Ints(jolts)
	ones, threes := 0, 1
	for i, n := range jolts {
		if i == 0 && n == 1 {
			ones++
		}
		if i > 0 {
			if (n - jolts[i-1]) == 1 {
				ones++
			} else if (n - jolts[i-1]) == 3 {
				threes++
			}
		}
	}
	fmt.Printf("1's: %d\n", ones)
	fmt.Printf("3's: %d\n", threes)
	fmt.Println(ones * threes)
}
