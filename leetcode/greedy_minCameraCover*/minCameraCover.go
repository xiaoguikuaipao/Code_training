package greedy_minCameraCover_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	ret := 0
	if traversal(root, &ret) == 0 {
		ret++
	}
	return ret
}

func traversal(node *TreeNode, ret *int) int {
	if node == nil {
		return 1
	}

	leftSon := traversal(node.Left, ret)
	rightSon := traversal(node.Right, ret)

	if leftSon == 0 || rightSon == 0 {
		*ret++
		return 2
	}
	if leftSon == 1 && rightSon == 1 {
		return 0
	}
	if leftSon == 2 || rightSon == 2 {
		return 1
	}
	return -1
}
