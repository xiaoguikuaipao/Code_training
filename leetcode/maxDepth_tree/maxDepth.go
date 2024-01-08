package maxDepth_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traversal(root, 0)
}

func traversal(node *TreeNode, depth int) int {
	if node != nil {
		depth++
	} else {
		return depth
	}
	maxDepth := depth
	leftDepth := traversal(node.Left, depth)
	rightDepth := traversal(node.Right, depth)
	if leftDepth > rightDepth {
		maxDepth = leftDepth
	} else {
		maxDepth = rightDepth
	}
	return maxDepth
}
