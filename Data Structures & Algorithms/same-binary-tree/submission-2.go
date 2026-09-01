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


		if pair.p == nil && pair.q == nil {
			continue
		}

		if pair.p == nil || pair.q == nil {
			return false
		}

		if pair.p.Val != pair.q.Val {
			return false
		}

		stack = append(stack, 
			Pair{pair.p.Left, pair.q.Left},
			Pair{pair.p.Right, pair.q.Right},
		)
	}

	return true
}
