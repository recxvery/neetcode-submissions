/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type State struct {
	Node *TreeNode
	Pos float64
	Neg float64
}

func isValidBST(root *TreeNode) bool {
	stack := []State{{root, math.Inf(1), math.Inf(-1)}}

	for len(stack) > 0 {
		state := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if state.Node == nil {
			continue
		}

		if float64(state.Node.Val) <= state.Neg || float64(state.Node.Val) >= state.Pos {
			return false
		}

		stack = append(stack,
			State{
				state.Node.Left,
				float64(state.Node.Val),
				state.Neg,
			},
			State{
				state.Node.Right,
				state.Pos,
				float64(state.Node.Val),
			},
		)
	}

	return true
}
