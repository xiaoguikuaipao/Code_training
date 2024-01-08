package reverseWords

import "slices"

func reverseWords(s string) string {
	bts := []byte(s)
	slices.Reverse(bts)
	size := len(bts)
	isLeftEdge, isRightEdge := false, false
	if bts[0] == ' ' {
		isLeftEdge = true
	}
	if bts[size-1] == ' ' {
		isRightEdge = true
	}
	for i, j := 0, size-1; i < j; {
		if bts[i] == ' ' && isLeftEdge == true {
			for bts[i] == ' ' {
				copy(bts[i:], bts[i+1:])
				size--
				j--
			}
			isLeftEdge = false
		}
		if bts[i] == ' ' && isLeftEdge == false {
			for bts[i] == ' ' && i+1 < size && bts[i+1] == ' ' {
				copy(bts[i:], bts[i+1:])
				size--
				j--
			}
			i++
		}
		if bts[j] == ' ' && isRightEdge == true {
			for bts[j] == ' ' {
				size--
				j--
			}
			isRightEdge = false
		}
		if bts[j] == ' ' && isRightEdge == false {
			for bts[j] == ' ' && j-1 > 0 && bts[j-1] == ' ' {
				copy(bts[j-1:], bts[j:])
				size--
				j--
			}
			j--
		}
		if bts[i] != ' ' && i < j {
			iRight := i
			for iRight < size-1 && bts[iRight+1] != ' ' {
				iRight++
			}
			for l, r := i, iRight; l < r; {
				bts[l], bts[r] = bts[r], bts[l]
				l++
				r--
			}
			i = iRight + 1
		}
		if bts[j] != ' ' && i < j {
			jLeft := j
			for jLeft > 0 && bts[jLeft-1] != ' ' {
				jLeft--
			}
			for l, r := jLeft, j; l < r; {
				bts[l], bts[r] = bts[r], bts[l]
				l++
				r--
			}
			j = jLeft - 1
		}

	}
	return string(bts[:size])
}
