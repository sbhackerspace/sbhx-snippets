// Steve Phillips / elimisteve
// 2014.01.03

package main

import "fmt"

type Node struct {
	Value       string
	Left, Right *Node
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
	preorder(tree)
	fmt.Printf("\n")
	inorder(tree)
	fmt.Printf("\n")
	postorder(tree)
	fmt.Printf("\n")
}

func preorder(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Printf("%s ", tree.Value)
	preorder(tree.Left)
	preorder(tree.Right)
}

func inorder(tree *Node) {
	if tree == nil {
		return
	}
	inorder(tree.Left)
	fmt.Printf("%s ", tree.Value)
	inorder(tree.Right)
}

func postorder(tree *Node) {
	if tree == nil {
		return
	}
	postorder(tree.Left)
	postorder(tree.Right)
	fmt.Printf("%s ", tree.Value)
}
