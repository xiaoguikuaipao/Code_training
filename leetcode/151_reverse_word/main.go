package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(reverseWords(" the  sky  is  blue "))
}

// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

// 示例 1：
// 输入：s = "the sky is blue"
// 输出："blue is sky the"

func reverseWords(s string) string {
	l := len(s)
	sb := []byte(s)
	i := 0
	for j := 0; j < l; j++ {
		if sb[j] != ' ' {
			if i == 0 {
				for j < l && sb[j] != ' ' {
					sb[i] = sb[j]
					i++
					j++
				}
			} else {
				sb[i] = ' '
				i++
				for j < l && sb[j] != ' ' {
					sb[i] = sb[j]
					i++
					j++
				}
			}
		}
	}
	slices.Reverse(sb[:i])
	for left, right := 0, 0; left < i && right < i; {
		for right < i && sb[right] != ' ' {
			right++
		}
		slices.Reverse(sb[left:right])
		right++
		left = right
	}
	return string(sb[:i])
}
