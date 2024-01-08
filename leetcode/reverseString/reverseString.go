package reverseString

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
