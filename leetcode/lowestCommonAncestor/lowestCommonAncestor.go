package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil && root != p && root != q {
		return nil
	} else if root == p || root == q {
		return root
	}

	if left == nil && right != nil {
		if root == p || root == q {
			return root
		} else {
			return right
		}
	}
	if left != nil && right == nil {
		if root == p || root == q {
			return root
		} else {
			return left
		}
	}

	if left != nil && right != nil {
		return root
	}

	return nil
}
