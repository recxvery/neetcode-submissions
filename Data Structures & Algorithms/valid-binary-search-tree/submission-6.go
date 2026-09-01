/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type State struct {
	node *TreeNode
	neg float64
	pos float64
}


func isValidBST(root *TreeNode) bool {
    stack := []State{{root, math.Inf(-1), math.Inf(1)}}

	for len(stack) > 0 {
		state := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if state.node == nil {
			continue
		}

		if float64(state.node.Val) >= state.pos || float64(state.node.Val) <= state.neg {
			return false
		}


		stack = append(stack, 
			State{
				state.node.Left,
				state.neg,
				float64(state.node.Val),
			},
			State {
				state.node.Right,
				float64(state.node.Val),
				state.pos,
			},
		)
	}

	return true
}
