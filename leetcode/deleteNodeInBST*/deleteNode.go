package deleteNodeInBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			leftSon := root.Left
			rightSon := root.Right
			root.Right = insertLeftToRight_Left(leftSon, rightSon)
			return root.Right
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func insertLeftToRight_Left(NodeToBeInserted, InsertingNode *TreeNode) *TreeNode {
	if InsertingNode.Left == nil {
		InsertingNode.Left = NodeToBeInserted
	} else {
		InsertingNode.Left = insertLeftToRight_Left(NodeToBeInserted, InsertingNode.Left)
	}
	return InsertingNode
}
