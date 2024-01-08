package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	leftSearch := searchBST(root.Left, val)
	if nil == leftSearch {
		return searchBST(root.Right, val)
	} else {
		return leftSearch
	}
}
