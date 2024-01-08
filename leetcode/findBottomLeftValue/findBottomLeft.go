package findBottomLeftValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeWithLabel struct {
	Node  *TreeNode
	Depth int
}

func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*NodeWithLabel, 0)
	queue = append(queue, &NodeWithLabel{root, 0})
	leftest := root.Val
	deepest := 0
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		if pop.Node.Left != nil {
			if pop.Depth == deepest {
				deepest++
				leftest = pop.Node.Left.Val
			}
			queue = append(queue, &NodeWithLabel{pop.Node.Left, pop.Depth + 1})
		}
		if pop.Node.Right != nil {
			if pop.Depth == deepest {
				deepest++
				leftest = pop.Node.Right.Val
			}
			queue = append(queue, &NodeWithLabel{pop.Node.Right, pop.Depth + 1})
		}
	}
	return leftest
}
