package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	res := subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 10000)
	if n >= 2 {
		res = dp(nums[:n/2], nums[n/2:])
	} else {
		res = append(res, []int{nums[0]}, []int{})
	}
	return res
}

func dp(nums1 []int, nums2 []int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 1 && n2 == 1 {
		res := make([][]int, 0)
		res = append(res, []int{}, []int{nums2[0]}, []int{nums1[0]}, []int{nums1[0], nums2[0]})
		return res
	} else if n1 == 1 {
		res := dp(nums2[:n2/2], nums2[n2/2:])
		populate := make([]int, 0)
		for _, s := range res {
			populate = append(s, nums1[0])
			res = append(res, populate)
		}
		return res
	} else if n2 == 1 {
		res := dp(nums1[:n1/2], nums1[n1/2:])
		populate := make([]int, 0)
		for _, s := range res {
			populate = append(s, nums2[0])
			res = append(res, populate)
		}
		return res
	} else {
		populate := make([][]int, 0)
		p1 := dp(nums1[:n1/2], nums1[n1/2:])
		p2 := dp(nums2[:n2/2], nums2[n2/2:])
		for _, e1 := range p1 {
			for _, e2 := range p2 {
				e3 := make([]int, len(e1)+len(e2))
				copy(e3, e1)
				copy(e3[len(e1):], e2)
				populate = append(populate, e3)
			}
		}
		for _, e := range populate {
			sort.Ints(e)
		}
		return populate
	}
}
