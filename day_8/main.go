package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
	"strings"
)

type (
	InstructionSet struct {
		instructions []*Instruction
		accumulator  int
	}
	Instruction struct {
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
	challenge1()
	challenge2()
}

func challenge1() {
	var instructions []*Instruction
	util.ForEachLine("day_8/input.txt", func(str string) {
		instructions = append(instructions, newInstruction(str))
	})
	set := InstructionSet{
		instructions: instructions,
	}
	fmt.Println(set.execute(0))
}

func (set *InstructionSet) execute(i int) (int, bool) {
	if i >= len(set.instructions) {
		return set.accumulator, true
	}
	inst := set.instructions[i]
	if inst.executed {
		return set.accumulator, false
	}
	inst.executed = true
	switch inst.op {
	case jmp:
		return set.execute(i+inst.val)
	case acc:
		set.accumulator += inst.val
		return set.execute(i+1)
	default:
		return set.execute(i+1)
	}
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

func newInstruction(str string) *Instruction {
	parts := strings.Split(str, " ")
	n, _ := strconv.Atoi(parts[1])
	return &Instruction{
		op:       newOp(parts[0]),
		val:      n,
		executed: false,
	}
}
