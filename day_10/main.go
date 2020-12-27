package main

import (
	"advent_2020/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//challenge1()
	challenge2()
}

func challenge2() {
	jolts := []int{0}
	util.ForEachLine("day_10/input.txt", func(str string) {
		n, _ := strconv.Atoi(str)
		jolts = append(jolts, n)
	})
	sort.Ints(jolts)
	combos := make(map[int]int)
	combos[0] = 1
	for i, j := range jolts {
		if i < (len(jolts)-1) && jolts[i+1] <= j+3 {
			combos[i+1] += combos[i]
		}
		if i < (len(jolts)-2) && jolts[i+2] <= j+3 {
			combos[i+2] += combos[i]
		}
		if i < (len(jolts)-3) && jolts[i+3] <= j+3 {
			combos[i+3] += combos[i]
		}
	}
	fmt.Println(combos[len(jolts)-1])
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
