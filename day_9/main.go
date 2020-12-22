package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
)

const cypherLength = 25

func main() {
	i := 0
	nums := make([]int, cypherLength)
	util.ForEachLine("day_9/input.txt", func(str string) {
		n, _ := strconv.Atoi(str)
		if i >= cypherLength {
			if !cubedSolution(nums, n) {
				fmt.Printf("Number doens't work: %d\n", n)
			}
		}
		nums[(i % cypherLength)] = n
		i++
	})
	fmt.Println("all numbers match")
}

func cubedSolution(nums []int, sum int) bool {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n+nums[j] == sum {
				return true
			}
		}
	}
	return false
}
