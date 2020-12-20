package main

type Node struct {
	counted  bool
	color    string
	children map[*Node]int
	parents  []*Node
}

func newNode(color string) *Node {
	return &Node{
		color:    color,
		children: make(map[*Node]int),
	}
}

type NodeMap struct {
	nodes map[string]*Node
}

func (n *NodeMap) getContainedCount(color string) int {
	c := 0
	for _, node := range n.nodes[color].parents {
		if r, ok := n.nodes[node.color]; ok && !r.counted {
			c++
			r.counted = true
		}
		c += n.getContainedCount(node.color)
	}
	return c
}

func (n *NodeMap) getContainingCount(color string) int {
	c := 0
	if rule, ok := n.nodes[color]; ok {
		for node, i := range rule.children {
			c += i
			c += i * n.getContainingCount(node.color)
		}
	}
	return c
}

func (n *NodeMap) getOrAdd(color string) *Node {
	node, ok := n.nodes[color]
	if ok {
		return node
	}
	n2 := newNode(color)
	n.nodes[color] = n2
	return n2
}