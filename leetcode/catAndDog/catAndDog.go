package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	m := 0
	fmt.Scan(&m)
	dogs := make([]dog, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&dogs[i].b, &dogs[i].c)
	}

	game(a, dogs)
}

func game(cats []int64, dogs []dog) {
	var sum int64
	for _, cat := range cats {
		sum += cat
	}
	for _, dogi := range dogs {
		var min int64 = 10000000
		var flag = false
		for _, cat := range cats {
			var consume1 int64
			var consume2 int64
			if cat > dogi.b && sum-cat > dogi.c {
				fmt.Println(0)
				flag = true
				break
			}
			if cat > dogi.b {
				consume1 = 0
			} else {
				consume1 = dogi.b - cat
			}
			if sum-cat > dogi.c {
				consume2 = 0
			} else {
				consume2 = dogi.c - sum + cat
			}
			if consume1+consume2 < min {
				min = consume1 + consume2
			}
		}
		if flag == false {
			fmt.Println(min)
		}
	}
}

type dog struct {
	b int64
	c int64
}
