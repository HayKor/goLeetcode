package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func areIdentical(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot // Check for both of them being nil
	}
	return root.Val == subRoot.Val &&
		areIdentical(root.Left, subRoot.Left) &&
		areIdentical(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var flag bool

	var walkTree func(root *TreeNode)
	walkTree = func(root *TreeNode) {
		if root == nil {
			return
		}

		flag = flag || areIdentical(root, subRoot)
		flag = flag || areIdentical(root.Left, subRoot)
		flag = flag || areIdentical(root.Right, subRoot)

		walkTree(root.Right)
		walkTree(root.Left)
	}

	walkTree(root)

	return flag
}

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	subRoot := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(isSubtree(&root, &subRoot))
}
