package greedy_candy

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	sum := 0
	for i := range candies {
		candies[i] = 1
		sum++
	}
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
			sum += candies[i-1] - candies[i] + 1
			candies[i] = candies[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			sum += candies[i+1] - candies[i] + 1
			candies[i] = candies[i+1] + 1
		}
	}
	return sum
}
