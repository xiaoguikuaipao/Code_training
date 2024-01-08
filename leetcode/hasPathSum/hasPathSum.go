package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	return dfs(root, sum, targetSum)
}

func dfs(node *TreeNode, sum, target int) bool {
	if node == nil {
		return false
	}
	if node.Val+sum == target {
		if node.Left == nil && node.Right == nil {
			return true
		}
	}
	leftSon := dfs(node.Left, node.Val+sum, target)
	RightSon := dfs(node.Right, node.Val+sum, target)
	return leftSon || RightSon
}
