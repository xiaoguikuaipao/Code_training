package constructMaximunBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	index, x := sliceMax(nums)
	root.Val = x
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}

func sliceMax(nums []int) (index, x int) {
	x = nums[0]
	for i := 0; i < len(nums); i++ {
		prex := x
		x = max(x, nums[i])
		if prex != x {
			index = i
		}
	}
	return index, x
}
