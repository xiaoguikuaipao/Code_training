package invertTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		if pop == nil {
			continue
		}
		pop.Left, pop.Right = pop.Right, pop.Left
		queue = append(queue, pop.Left)
		queue = append(queue, pop.Right)
	}
	return root
}
