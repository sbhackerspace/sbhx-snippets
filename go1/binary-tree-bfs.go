// Steve Phillips / elimisteve
// 2014.01.03

package main

import "fmt"

type Node struct {
	Value       string
	Left, Right *Node
}

func (node *Node) Children() []*Node {
	var c []*Node
	if node.Left != nil {
		c = append(c, node.Left)
	}
	if node.Right != nil {
		c = append(c, node.Right)
	}
	return c
}

func BFS(node *Node, value string) *Node {
	q := []*Node{node}
	var n *Node
	for len(q) > 0 {
		n, q = q[0], q[1:] // Pop
		fmt.Printf("%s ", n.Value)
		if n.Value == value {
			return n
		}
		q = append(q, n.Children()...)
	}
	return nil
}

// See https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Sorted_binary_tree.svg/333px-Sorted_binary_tree.svg.png
var tree = &Node{
	Value: "F",
	Left: &Node{
		Value: "B",
		Left:  &Node{Value: "A"},
		Right: &Node{
			Value: "D",
			Left:  &Node{Value: "C"},
			Right: &Node{Value: "E"},
		},
	},
	Right: &Node{
		Value: "G",
		Right: &Node{
			Value: "I",
			Left:  &Node{Value: "H"},
		},
	},
}

func main() {
	if n := BFS(tree, "I"); n != nil {
		fmt.Printf("\nFound %s\n", n.Value)
	}
	fmt.Printf("\n")
	if n := BFS(tree, "H"); n != nil {
		fmt.Printf("\nFound %s\n", n.Value)
	}
}
