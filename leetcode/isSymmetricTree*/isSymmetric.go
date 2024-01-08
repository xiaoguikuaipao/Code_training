package isSymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	result := false
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		result = outside && inside
	}
	return result
}
