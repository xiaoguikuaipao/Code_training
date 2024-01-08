package countNodes_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var leftCount, rightCount int
	if root == nil {
		return 0
	}
	leftCount = countNodes(root.Left)
	rightCount = countNodes(root.Right)
	return leftCount + rightCount + 1
}
