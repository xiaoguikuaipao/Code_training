package getMinimunDifference

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minD := math.MaxInt32
	getMin(root, &minD)
	return minD
}

func getMin(node *TreeNode, minD *int) (minNode, maxNode *TreeNode) {

	if node == nil {
		return nil, nil
	}

	leftMin, leftMax := getMin(node.Left, minD)
	rightMin, rightMax := getMin(node.Right, minD)
	if leftMax != nil && rightMin != nil && int(math.Abs(float64(leftMax.Val-rightMin.Val))) < *minD {
		*minD = int(math.Abs(float64(leftMax.Val - rightMin.Val)))
	}
	if leftMax != nil && int(math.Abs(float64(leftMax.Val-node.Val))) < *minD {
		*minD = int(math.Abs(float64(leftMax.Val - node.Val)))
	}
	if rightMin != nil && int(math.Abs(float64(node.Val-rightMin.Val))) < *minD {
		*minD = int(math.Abs(float64(node.Val - rightMin.Val)))
	}
	maxNode = node
	minNode = node
	if leftMax != nil && leftMax.Val >= node.Val {
		return nil, nil
	}
	if rightMin != nil && rightMin.Val <= node.Val {
		return nil, nil
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
	return minNode, maxNode
}
