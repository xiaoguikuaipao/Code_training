package main

import "fmt"

func main() {
	nums1 := []int{0, 0}
	m := 0
	nums2 := []int{2, 5}
	n := 2
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	place := m + n - 1
	f1 := m - 1
	f2 := n - 1
	if m == 0 {
		copy(nums1, nums2)
	}
	for f1 >= 0 && f2 >= 0 {
		if nums1[f1] > nums2[f2] {
			nums1[place] = nums1[f1]
			f1--
			place--
		} else {
			nums1[place] = nums2[f2]
			f2--
			place--
		}
	}
	for place >= 0 {
		if f1 >= 0 {
			nums1[place] = nums1[f1]
			place--
			f1--
		}
		if f2 >= 0 {
			nums1[place] = nums2[f2]
			place--
			f2--
		}
	}
}
