package main

import (
	"fmt"
)

func main() {
	// create an example max depth tree
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// use the max depth method to find the max depth of the tree
	s := Solution{}
	fmt.Println(s.MaxDepth(root)) // Output: 3

}
