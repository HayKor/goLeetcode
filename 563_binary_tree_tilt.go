package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sumOfNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sumOfNode(root.Left) + sumOfNode(root.Right)
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	absDiff := sumOfNode(root.Left) - sumOfNode(root.Right)
	if absDiff < 0 {
		absDiff = -absDiff
	}
	return absDiff + findTilt(root.Right) + findTilt(root.Left)
}

// func main() {
// 	longRoot := TreeNode{
// 		Val: 4,
// 		Left: &TreeNode{
// 			Val:   2,
// 			Left:  &TreeNode{Val: 3},
// 			Right: &TreeNode{Val: 5},
// 		},
// 		Right: &TreeNode{
// 			Val:   9,
// 			Right: &TreeNode{Val: 7},
// 		},
// 	}
// 	fmt.Println(findTilt(&longRoot))
// }
