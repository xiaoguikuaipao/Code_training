package main

// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
func removeDuplicates(nums []int) int {
	l := len(nums)
	var ret int
	for i, j := 0, 0; j < l; {
		for j < l && nums[i] == nums[j] {
			j++
		}
		if j < l {
			nums[i+1] = nums[j]
			i++
		}
		ret = i
	}
	return ret + 1
}

func main()
