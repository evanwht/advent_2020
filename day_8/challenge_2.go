package main

import (
	"advent_2020/util"
	"fmt"
	"strconv"
	"strings"
)

type (
	InstructionNode struct {
		index           int
		op              operation
		value           int
		parents         []*InstructionNode
		possibleParents []*InstructionNode
	}
	InstructionMap struct {
		nodes map[int]*InstructionNode
	}
)

func challenge2() {
	nodes := InstructionMap{
		nodes: make(map[int]*InstructionNode),
	}
	i := 0
	util.ForEachLine("day_8/input.txt", func(str string) {
		node := nodes.getOrAdd(i, str)
		if node.op == jmp {
			n := nodes.getOrAdd(i + node.value)
			n.parents = append(n.parents, node)
			n = nodes.getOrAdd(i + 1)
			n.possibleParents = append(n.possibleParents, node)
		} else if node.op == nop {
			n := nodes.getOrAdd(i + 1)
			n.parents = append(n.parents, node)
			n = nodes.getOrAdd(i + node.value)
			n.possibleParents = append(n.possibleParents, node)
		}
		i++
	})
	// walk back from end node to see what node to change
	fmt.Println(nodes.walk(len(nodes.nodes) - 1))
}

func (m *InstructionMap) walk(i int) *InstructionNode {
	node := m.nodes[i]
	for _, n := range node.parents {
		node = m.walk(n.index)
		if len(node.possibleParents) > 0 {
			return node.possibleParents[0]
		}
	}
	if len(node.possibleParents) > 0 {
		return node.possibleParents[0]
	}
	return node
}

func newNode(str string) *InstructionNode {
	parts := strings.Split(str, " ")
	n, _ := strconv.Atoi(parts[1])
	return &InstructionNode{
		op:      newOp(parts[0]),
		value:   n,
		parents: []*InstructionNode{},
	}
}

func (m *InstructionMap) getOrAdd(i int, str ...string) *InstructionNode {
	var n *InstructionNode
	if len(str) == 0 {
		n = &InstructionNode{
			parents: []*InstructionNode{},
		}
	} else {
		n = newNode(str[0])
	}
	n.index = i
	node, ok := m.nodes[i]
	if ok {
		n.parents = node.parents
	}
	m.nodes[i] = n
	return n
}
