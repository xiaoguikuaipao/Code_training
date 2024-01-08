package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	traversal_recur(root, &ret)
	return ret
}

func traversal_recur(node *TreeNode, pret *[]int) {
	if node == nil {
		return
	}
	//this is correct, but
	//ret := *pret

	//append may return a new slice which refers to a new basic arrays.
	*pret = append(*pret, node.Val)
	traversal_recur(node.Left, pret)
	traversal_recur(node.Right, pret)
}

func preorderTraversal_Notrecur(node *TreeNode, pret *[]int) {
	st := make([]*TreeNode, 0)
	st = append(st, node)
	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		if cur == nil {
			continue
		}

		*pret = append(*pret, cur.Val)
		st = append(st, cur.Right)
		st = append(st, cur.Left)
	}
}

// postorderTraversal_Notrecur modify the preorderTraversal to get the code
func postorderTraversal_Notrecur(node *TreeNode, pret *[]int) {
	st := make([]*TreeNode, 0)
	st = append(st, node)
	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		if cur == nil {
			continue
		}

		*pret = append(*pret, cur.Val)
		st = append(st, cur.Left)
		st = append(st, cur.Right)
	}
	size := len(*pret)
	for i, j := 0, size-1; i < j; {
		(*pret)[i], (*pret)[j] = (*pret)[j], (*pret)[i]
		i++
		j--
	}
}

// inorderTraversal_Notrecur is different from post and pre, because the access and process are not simultaneous
func inorderTraversal_Notrecur(node *TreeNode, pret *[]int) {

	//use stack and index
	it := node
	st := make([]*TreeNode, 0)
	for it != nil || len(st) > 0 {

		if it != nil {
			st = append(st, it)
			it = it.Left
		} else {
			it = st[len(st)-1]
			st = st[:len(st)-1]
			*pret = append(*pret, it.Val)
			it = it.Right
		}
	}
}
