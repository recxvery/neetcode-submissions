/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	lvl := 0
	for len(queue) > 0 {
		lvl_size := len(queue)
		tmp := []int{}
		
		for range lvl_size {
			if lvl % 2 == 0 {
				node := queue[0]
				queue[0] = nil
				queue = queue[1:]

				tmp = append(tmp, node.Val)

				if node.Left != nil {
					queue = append(queue, node.Left)
				}

				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				node := queue[len(queue) - 1]
				queue[len(queue) - 1] = nil
				queue = queue[:len(queue) - 1]

				tmp = append(tmp, node.Val)

				if node.Right != nil {
					queue = append([]*TreeNode{node.Right}, queue...)
				}

				if node.Left != nil {
					queue = append([]*TreeNode{node.Left}, queue...)
				}
			}
		}

		lvl++
		res = append(res, tmp)
	}

	return res
}

