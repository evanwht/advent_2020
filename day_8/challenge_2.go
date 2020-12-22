package main

import (
	"advent_2020/util"
	"fmt"
)

func challenge2() {
	var nodes []*Instruction
	util.ForEachLine("day_8/input.txt", func(str string) {
		nodes = append(nodes, newInstruction(str))
	})
	for i, n := range nodes {
		if n.op == jmp || n.op == nop {
			nodes[i] = n.swap()
			if num, ok := works(nodes); ok {
				fmt.Println(num)
				return
			}
			nodes[i] = nodes[i].swap()
		}
	}
	fmt.Println("nothing")
}

func works(nodes []*Instruction) (int, bool) {
	set := InstructionSet{
		instructions: nodes,
	}
	num, ok := set.execute(0)
	if !ok {
		for _, n := range nodes {
			n.executed = false
		}
	}
	return num, ok
}

func (n *Instruction) swap() *Instruction {
	if n.op == jmp {
		return &Instruction {
			op: nop,
			val: n.val,
			executed: n.executed,
		}
	} else if n.op == nop {
		return &Instruction {
			op: jmp,
			val: n.val,
			executed: n.executed,
		}
	}
	return n
}
