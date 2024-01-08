package minDepth_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	ret := depth
	leftDepth := dfs(node.Left, depth)
	rightDepth := dfs(node.Right, depth)
	if leftDepth == depth {
		return rightDepth
	}
	if rightDepth == depth {
		return leftDepth
	}
	if leftDepth > rightDepth {
		ret = rightDepth
	} else {
		ret = leftDepth
	}
	return ret
}
