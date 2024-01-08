package BSTToGST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// right middle left --- traversal
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	return RMLTraversal(root, &pre)
}

// RMLTraversal Use double pointer
func RMLTraversal(root *TreeNode, pre *int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right = RMLTraversal(root.Right, pre)
	root.Val = *pre + root.Val
	*pre = root.Val
	root.Left = RMLTraversal(root.Left, pre)
	return root
}
