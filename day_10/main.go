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
	var jolts []int
	util.ForEachLine("day_10/test_input.txt", func(str string) {
		n, _ := strconv.Atoi(str)
		jolts = append(jolts, n)
	})
	sort.Ints(jolts)
	combos := 1
	for i := 0; i < len(jolts); i++ {
		m, k := 0, 1
		for (k+i) < len(jolts) && (jolts[k+i]-jolts[i]) <= 3 {
			m++
			k++
		}
		if m > 0 {
			combos *= m
		}
	}
	fmt.Println(combos)
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
