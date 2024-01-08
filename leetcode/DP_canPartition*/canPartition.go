package main

func canPartition(nums []int) bool {
	partA := 0
	partB := 0
	for _, e := range nums {
		partB += e
	}

	return backtracking(partA, partB, nums, 0)
}

func backtracking(partA int, partB int, nums []int, start int) bool {
	if partA == partB {
		return true
	}
	for i := start; i < len(nums); i++ {
		partA += nums[i]
		partB -= nums[i]

		if backtracking(partA, partB, nums, i+1) {
			return true
		}

		partA -= nums[i]
		partB += nums[i]
	}
	return false
}
