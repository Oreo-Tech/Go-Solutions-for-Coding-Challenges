package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode) *TreeNode {

	fmt.Println("Tree is", root)
	if root.Left != nil {
		PrintTree(root.Left)
	}
	if root.Right != nil {
		PrintTree(root.Right)
	}
	return nil
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}

func main() {

	t := &TreeNode{Val: 4, Left: nil, Right: nil}
	t.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	t.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	t.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	t.Left.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	t.Right.Left = &TreeNode{Val: 6, Left: nil, Right: nil}
	t.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}
	fmt.Println("Before")
	PrintTree(t)
	//fmt.Println("Tree root is", t)
	t = invertTree(t)
	//fmt.Println("Tree root is", t)
	fmt.Println("After")
	PrintTree(t)
}
