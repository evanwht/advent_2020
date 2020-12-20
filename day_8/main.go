package main

import (
	"advent_2020/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	instructionSet struct {
		instructions []*instruction
		accumulator  int
	}
	instruction struct {
		op       operation
		val      int
		executed bool
	}
	operation string
)

const (
	acc operation = "acc"
	jmp operation = "jmp"
	nop operation = "nop"
)

func main() {
	file, err := os.Open("day_8/input.txt")
	util.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []*instruction
	for scanner.Scan() {
		instructions = append(instructions, newInstruction(scanner.Text()))
	}
	set := instructionSet{
		instructions: instructions,
	}
	fmt.Println(set.execute(0))
}

func (set *instructionSet) execute(i int) int {
	if i >= len(set.instructions) {
		return set.accumulator
	}
	inst := set.instructions[i]
	if inst.executed {
		return set.accumulator
	}
	inst.executed = true
	switch inst.op {
	case jmp:
		set.execute(i + inst.val)
	case acc:
		set.accumulator += inst.val
		set.execute(i + 1)
	default:
		set.execute(i + 1)
	}
	return set.accumulator
}

func newOp(str string) operation {
	switch str {
	case "acc":
		return acc
	case "jmp":
		return jmp
	default:
		return nop
	}
}

func newInstruction(str string) *instruction {
	parts := strings.Split(str, " ")
	n, _ := strconv.Atoi(parts[1])
	return &instruction{
		op:       newOp(parts[0]),
		val:      n,
		executed: false,
	}
}
