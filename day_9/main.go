package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
)

const cypherLength = 25
const badNumber = 167829540

func main() {
	//challenge1()
	challenge2()
}

type Sum struct {
	start int
	end   int
	sum   int
}

func challenge2() {
	var nums []int
	var sums []*Sum
	util.ForEachLine("day_9/input.txt", func(str string) {
		n, _ := strconv.Atoi(str)
		nums = append(nums, n)
	})
	for i, n := range nums {
		sums = append(sums, &Sum{
			start: i,
			end: i,
		})
		for _, s := range sums {
			s.sum += n
			s.end++
			for s.sum > badNumber && s.start < s.end {
				s.sum -= nums[s.start]
				s.start++
			}
			if s.sum == badNumber && (s.end - s.start) > 1{
				fmt.Println(min(nums[s.start:s.end]) + max(nums[s.start:s.end]))
				return
			}
		}
		i++
	}
	fmt.Println("all numbers match")
}

func min(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func challenge1() {
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
