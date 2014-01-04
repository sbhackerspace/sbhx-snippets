// Steve Phillips / elimisteve
// 2014.01.03

package main

import "fmt"

type Node struct {
	Value       string
	Left, Right *Node
}

func (node *Node) DFS(value string) *Node {
	if node == nil {
		return nil
	}
	fmt.Printf("%s ", node.Value)
	if node.Value == value {
		return node
	}
	if n := node.Left.DFS(value); n != nil {
		return n
	}
	if n := node.Right.DFS(value); n != nil {
		return n
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
	if n := tree.DFS("C"); n != nil {
		fmt.Printf("\nFound %s\n\n", n.Value)
	}
	if n := tree.DFS("I"); n != nil {
		fmt.Printf("\nFound %s\n", n.Value)
	}
}
