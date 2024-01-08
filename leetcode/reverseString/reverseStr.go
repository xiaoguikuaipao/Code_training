package reverseString

import "sync"

func reverseStr(s string, k int) string {
	var wg sync.WaitGroup
	bts := []byte(s)
	index := 0
	count := 0
	size := len(s)
	for index < size {
		count++
		if count == 2*k {
			count = 0
			left := index + 1 - 2*k
			right := left + k - 1
			wg.Add(1)
			go func() {
				for left < right {
					bts[left], bts[right] = bts[right], bts[left]
					left++
					right--
				}
				wg.Done()
			}()
		}
		index++
	}

	if count < k {
		left := 0
		if index > k {
			left = index - count
		}
		right := left + count - 1
		for left < right {
			bts[left], bts[right] = bts[right], bts[left]
			left++
			right--
		}
	}

	if count >= k {
		left := 0
		if index > k {
			left = index - count
		}
		right := left + k - 1
		for left < right {
			bts[left], bts[right] = bts[right], bts[left]
			left++
			right--
		}
	}
	wg.Wait()
	return string(bts)
}
