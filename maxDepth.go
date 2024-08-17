package main

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution is a struct to hold the solution method.
type Solution struct{}

func (s *Solution) MaxDepth(root *TreeNode) int {
	// handle the case in which there are no nodes
	if root == nil {
		return 0
	}

	// check the depths of both sides of the tree
	lDepth := s.MaxDepth(root.Left)
	rDepth := s.MaxDepth(root.Right)

	// check if left is bigger
	if lDepth > rDepth {
		return lDepth + 1 // the +1 is to account for the current node
	}
	return rDepth + 1
}
