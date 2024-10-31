package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func areIdentical(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot // Check for both of them being nil
	}

	if root.Val == subRoot.Val {
		return areIdentical(root.Left, subRoot.Left) && areIdentical(root.Right, subRoot.Right)
	} else {
		return false
	}
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if areIdentical(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}

// func main() {
// 	root := TreeNode{
// 		Val: 3,
// 		Left: &TreeNode{
// 			Val:  4,
// 			Left: &TreeNode{Val: 1},
// 			Right: &TreeNode{ // the subRoot
// 				Val: 2,
// 				Left: &TreeNode{
// 					Val: 0,
// 				},
// 			},
// 		},
// 		Right: &TreeNode{
// 			Val: 5,
// 		},
// 	}
// 	subRoot := TreeNode{
// 		Val: 2,
// 		Left: &TreeNode{
// 			Val: 0,
// 		},
// 	}
// 	fmt.Println(isSubtree(&root, &subRoot))
// }
