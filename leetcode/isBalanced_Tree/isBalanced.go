package isBalanced_Tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if isBalanced(root.Left) && isBalanced(root.Right) {
		left := getHeight(root.Left)
		right := getHeight(root.Right)
		if math.Abs(float64(left-right)) <= 1 {
			return true
		} else {
			return false
		}
	}
	return false
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	var maxHeight int
	if leftHeight > rightHeight {
		maxHeight = leftHeight
	} else {
		maxHeight = rightHeight
	}
	return maxHeight + 1
}
