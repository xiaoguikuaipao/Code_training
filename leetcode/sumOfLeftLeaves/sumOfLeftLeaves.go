package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return postOrder(root)
}

func postOrder(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		leftSon := node.Left.Val
		return leftSon + postOrder(node.Right)
	} else {
		return postOrder(node.Left) + postOrder(node.Right)
	}
}
