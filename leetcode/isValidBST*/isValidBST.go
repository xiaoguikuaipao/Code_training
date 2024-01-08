package isValidBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	_, _, B := isValidBSTWithM(root)
	return B
}

func isValidBSTWithM(node *TreeNode) (minNode, maxNode *TreeNode, B bool) {
	if node == nil {
		return nil, nil, true
	}

	leftMin, leftMax, leftB := isValidBSTWithM(node.Left)
	rightMin, rightMax, rightB := isValidBSTWithM(node.Right)
	if leftB && rightB {
		maxNode = node
		minNode = node
		if leftMax != nil && leftMax.Val >= node.Val {
			return nil, nil, false
		}
		if rightMin != nil && rightMin.Val <= node.Val {
			return nil, nil, false
		}
		if leftMax != nil && leftMax.Val > maxNode.Val {
			maxNode = leftMax
		}
		if rightMax != nil && rightMax.Val > maxNode.Val {
			maxNode = rightMax
		}
		if leftMin != nil && leftMin.Val < minNode.Val {
			minNode = leftMin
		}
		if rightMin != nil && rightMin.Val < minNode.Val {
			minNode = rightMin
		}
		return minNode, maxNode, true
	} else {
		return nil, nil, false
	}

}
