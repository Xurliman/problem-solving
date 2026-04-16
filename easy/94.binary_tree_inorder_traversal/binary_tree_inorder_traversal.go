package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	recurse := func(root *TreeNode) {}
	recurse = func(root *TreeNode) {
		if root == nil {
			return
		}
		recurse(root.Left)
		res = append(res, root.Val)
		recurse(root.Right)
	}

	recurse(root)
	return res
}
