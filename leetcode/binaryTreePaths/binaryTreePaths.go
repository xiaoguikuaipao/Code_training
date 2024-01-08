package binaryTreePaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	path := make([]byte, 0)
	ret := make([]string, 0)
	path = append(path, []byte(strconv.Itoa(root.Val))...)
	dfs(root, &path, &ret)
	return ret
}

func dfs(node *TreeNode, path *[]byte, ret *[]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*ret = append(*ret, string(*path))
		return
	}
	if node.Left != nil {
		leftPath := []byte("->" + strconv.Itoa(node.Left.Val))
		*path = append(*path, leftPath...)
		dfs(node.Left, path, ret)
		*path = (*path)[:(len(*path) - len(leftPath))]
	}
	if node.Right != nil {
		rightPath := []byte("->" + strconv.Itoa(node.Right.Val))
		*path = append(*path, rightPath...)
		dfs(node.Right, path, ret)
		*path = (*path)[:(len(*path) - len(rightPath))]
	}
}
