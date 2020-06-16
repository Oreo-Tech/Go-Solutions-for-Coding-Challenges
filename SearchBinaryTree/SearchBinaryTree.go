package main

import "fmt"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
//  func searchBST(root *TreeNode, val int) *TreeNode {
// 	 fmt.Println("Inside root is",root)
//     if root==nil || root.Val==val {
// 		fmt.Println("Success root is",root)
// 		return root
		
// 	}
// 	if(val>root.Val) {
// 		root=root.Right
// 	}else{
// 		root=root.Left
// 	}
// 	return searchBST(root,val)
// }

func searchBST(root *TreeNode, val int) *TreeNode {


for root !=nil{
	if root.Val==val{
		return root
	}else if root.Left.Val==val{
		return root.Left
	}else if root.Right.Val==val{
		return root.Right
	}else if val>root.Val{
		root=root.Right	
	}else{
		root=root.Left
	}
}
	return root
}
func main() {
	root:=&TreeNode{Val: 4,Left:  nil,Right:  nil}
	root.Left=&TreeNode{Val: 2,Left: &TreeNode{Val: 1},Right: &TreeNode{Val: 7}}
	root.Right=&TreeNode{ 7,nil,nil}
	result:=searchBST(root,2)
	fmt.Println("Root",result)

}