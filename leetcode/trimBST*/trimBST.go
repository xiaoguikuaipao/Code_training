package trimBST_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low || root.Val > high {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil && root.Right != nil {
			return trimBST(root.Right, low, high)
		} else if root.Left != nil && root.Right == nil {
			return trimBST(root.Left, low, high)
		} else {
			root.Right = mergeLToR(root.Left, root.Right)
			root.Right = trimBST(root.Right, low, high)
			return root.Right
		}
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func mergeLToR(NodeToBeInserted *TreeNode, InsertingNode *TreeNode) *TreeNode {
	if InsertingNode.Left == nil {
		InsertingNode.Left = NodeToBeInserted
	} else {
		InsertingNode.Left = mergeLToR(NodeToBeInserted, InsertingNode.Left)
	}
	return InsertingNode
}
