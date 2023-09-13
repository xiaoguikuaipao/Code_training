package TreeSort

type tree struct {
	data       int
	leftChild  *tree
	rightChild *tree
}

func TreeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	output(values[:0], root)
}
func output(values []int, node *tree) []int {
	if node != nil {
		values = output(values, node.leftChild)
		values = append(values, node.data)
		values = output(values, node.rightChild)
	}
	return values
}

func add(node *tree, v int) *tree {
	if node == nil {
		node = new(tree)
		node.data = v
		return node
	} else {
		if node.data < v {
			node.rightChild = add(node.rightChild, v)
		}
		if node.data >= v {
			node.leftChild = add(node.leftChild, v)
		}
	}
	return node
}
