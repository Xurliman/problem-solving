package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := countHeight(root.Left), countHeight(root.Right)
	if left-right > 1 || left-right < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func countHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(countHeight(root.Left), countHeight(root.Right)) + 1
}
