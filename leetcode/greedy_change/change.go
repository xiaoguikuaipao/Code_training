package greedy_change

func lemonadeChange(bills []int) bool {
	money := make([]int, 3)
	for i := 0; i < len(bills); i++ {
		diff := 0
		if bills[i] == 5 {
			money[0]++
			continue
		} else if bills[i] == 10 {
			money[1]++
			diff = 5
			if money[0] > 0 {
				money[0]--
			} else {
				return false
			}
			continue
		} else if bills[i] == 20 {
			money[2]++
			diff = 15
			for diff > 0 {
				if money[1] > 0 && diff >= 10 {
					diff -= 10
					money[1]--
					continue
				}
				if money[0] > 0 {
					diff -= 5
					money[0]--
					continue
				}
				break
			}
			if diff != 0 {
				return false
			}
		}
	}
	return true
}
