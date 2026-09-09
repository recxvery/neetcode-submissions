/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	right_h := Depth(root.Right)
	left_h := Depth(root.Left)

	if abs(right_h, left_h) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)	
}

func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(Depth(root.Left), Depth(root.Right)) + 1
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}

	return a - b
}