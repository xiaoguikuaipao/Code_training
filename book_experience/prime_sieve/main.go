package main

import "fmt"

func generate(ch chan<- int) {
	for i := 2; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// 主程序i 代表输出素数个数
// func main() {
// 	ch := make(chan int)
// 	go generate(ch)
// 	for i := 0; i < 10; i++ {
// 		prime := <-ch
// 		fmt.Printf("get prime %d from ch %p\n", prime, ch)
// 		print(prime, "\n")
// 		ch1 := make(chan int)
// 		go Filter(ch, ch1, prime)
// 		ch = ch1
// 	}
// }

func NewFilter(in <-chan int, prime int) (out chan int) {
	fmt.Println("Filter ", prime, "has been created")
	out = make(chan int)
	go func() {
		defer fmt.Println("filter ", prime, "has been closed")
		for {
			i, ok := <-in
			if !ok {
				close(out)
				break
			}
			s := fmt.Sprintf("%d enter Fileter[%d]:", i, prime)
			if i%prime != 0 {
				out <- i
				s = fmt.Sprintf("%s [pass]", s)
			} else {

				s = fmt.Sprintf("%s [be filted]", s)
			}
			fmt.Println(s)
		}
	}()
	return out
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		print(prime, "\n")
		ch = NewFilter(ch, prime)
	}
}
