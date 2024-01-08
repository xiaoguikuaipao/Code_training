package main

func monotoneIncreasingDigits(n int) int {
	digits := make([]int, 0)
	breakUp(n, &digits)
	construct := make([]int, 0)
	sum := 0
	backtracking(&sum, digits, &construct, len(digits)-1, false)
	return sum
}

func backtracking(sum *int, digits []int, ret *[]int, index int, degrade bool) bool {
	if len(*ret) == len(digits) {
		factor := 1
		for i := len(*ret) - 1; i >= 0; i-- {
			*sum += (*ret)[i] * factor
			factor *= 10
		}
		return true
	}
	if index < 0 {
		return false
	}
	start := 0
	if degrade == false {
		start = digits[index]
	} else {
		start = 9
	}
	for j := start; j >= 0; j-- {
		if j != start {
			degrade = true
		}
		if len(*ret) == 0 || j >= (*ret)[len(*ret)-1] {
			*ret = append(*ret, j)
			if backtracking(sum, digits, ret, index-1, degrade) {
				return true
			}
			*ret = (*ret)[:len(*ret)-1]
		} else {
			break
		}
	}
	return false
}

func breakUp(n int, digits *[]int) {
	for n != 0 {
		remainder := n % 10
		*digits = append(*digits, remainder)
		n = n / 10
	}
}

func main() {
	monotoneIncreasingDigits(99998)
}
