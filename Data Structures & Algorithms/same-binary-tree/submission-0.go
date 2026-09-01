/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
	p *TreeNode
	q *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []Pair{{p, q}}

	for len(stack) > 0 {
		pair := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		p, q = pair.p, pair.q

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		stack = append(stack, 
			Pair{p.Left, q.Left},
			Pair{p.Right, q.Right},
		)
	}

	return true
}
