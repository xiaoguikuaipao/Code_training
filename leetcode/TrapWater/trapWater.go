package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := TrapWater(height)
	fmt.Println(res)
}

func TrapWater(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	water := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
