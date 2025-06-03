package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	st := NewSyncStack[string](100)
	var wg sync.WaitGroup
	rs := make(chan bool, 10000)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			ok := st.Push(fmt.Sprintf("goroutine %d in", i))
			rs <- ok
		}(i)
	}
	wg.Wait()
	close(rs)
	var ok, no int
	for r := range rs {
		if r {
			ok++
		} else {
			no++
		}
	}
	fmt.Println("ok", ":", ok, " ", "no", ":", no)
	fmt.Println(st.Len())
}
