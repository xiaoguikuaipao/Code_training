package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	quickSort(nums, left, right)
}

func quickSort(nums []int, left, right int) {
	i, j := left, right
	if i < j {
		temp := nums[left]
		for i < j {
			for i < j && nums[j] > temp {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			for i < j && nums[i] < temp {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = temp
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)
		for i < j && nums[j] > temp {
			j--
		}
		if i < j && nums[j] < temp {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] < temp {
			i++
		}
		if i < j && nums[i] > temp {
			nums[j] = nums[i]
			j--
		}
		nums[i] = temp
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)

	}
}
