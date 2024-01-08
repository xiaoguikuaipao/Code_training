package main

import (
	"fmt"
	"sync"
)

// arrays - []string
// get the max all urls

func getSumOfUrls(urls []string) int {
	result := 0
	maxGoroutines := 20
	control := make(chan struct{}, maxGoroutines)
	collect := make(chan int, 0)
	endSemaphore := make(chan struct{}, 0)
	var wg sync.WaitGroup

	go func() {
		for e := range collect {
			result += e
		}
		endSemaphore <- struct{}{}
	}()

	for _, url := range urls {
		wg.Add(1)
		control <- struct{}{}
		go func(url string) {
			defer func() {
				<-control
				wg.Done()
			}()
			size := getSize(url)
			collect <- size
		}(url)
	}
	wg.Wait()
	close(collect)

	<-endSemaphore

	return result
}

func getSize(url string) int {
	return 1
}

func main() {
	fmt.Println(getSumOfUrls([]string{"1", "2", "3"}))
}
