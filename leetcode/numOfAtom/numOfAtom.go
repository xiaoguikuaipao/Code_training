package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(countOfAtoms("(B2O39He17BeBe49)39"))
}

func countOfAtoms(formula string) string {
	table := make(map[string]int)
	restack := make([]string, 0)
	stack := make([]string, 0)
	for i, size := 0, len(formula); i < size; i++ {
		//如果是原子，入栈，初始形式为O，S, Mg
		if formula[i] >= 'A' && formula[i] <= 'Z' {
			if i < size-1 && formula[i+1] >= 'a' && formula[i+1] <= 'z' {
				stack = append(stack, formula[i:i+2])
				i = i + 1
				continue
			} else {
				stack = append(stack, formula[i:i+1])
			}
		}
		// 如果是括号直接入栈
		if formula[i] == '(' || formula[i] == ')' {
			stack = append(stack, formula[i:i+1])
		}
		//如果是数字，乘
		if unicode.IsNumber(rune(formula[i])) {
			start := i
			size := len(formula)
			for i < size {
				if i < size-1 && unicode.IsNumber(rune(formula[i+1])) {
					i++
				} else {
					break
				}
			}
			// 如果是()后的数字
			if stack[len(stack)-1] == ")" {
				stack = stack[:len(stack)-1]
				for {
					if unicode.IsNumber(rune(stack[len(stack)-1][len(stack[len(stack)-1])-1])) {
						digit := 1
						for count := len(stack[len(stack)-1]) - 2; count > 0; count-- {
							if unicode.IsNumber(rune(stack[len(stack)-1][count])) {
								digit++
							} else {
								break
							}
						}
						now, _ := strconv.Atoi(string(stack[len(stack)-1][len(stack[len(stack)-1])-digit : len(stack[len(stack)-1])]))
						current, _ := strconv.Atoi(string(formula[start : i+1]))
						now = current * now
						restack = append(restack, stack[len(stack)-1][:len(stack[len(stack)-1])-digit]+strconv.Itoa(now))
						stack = stack[:len(stack)-1]
					} else {
						current, _ := strconv.Atoi(string(formula[start : i+1]))
						restack = append(restack, stack[len(stack)-1][:len(stack[len(stack)-1])]+strconv.Itoa(current))
						stack = stack[:len(stack)-1]
					}
					if stack[len(stack)-1] == "(" {
						//把( 弹出，然后把乘后的原子放回
						stack = stack[:len(stack)-1]
						for len(restack) > 0 {
							stack = append(stack, restack[len(restack)-1])
							restack = restack[:len(restack)-1]
						}
						break
					}
				}
			} else {
				//将(后的每个原子乘上数字放回栈中，形式为O3, S3
				// 如果数字是 ()中的
				top := len(stack) - 1
				current, _ := strconv.Atoi(string(formula[start : i+1]))
				stack[top] = stack[top][:len(stack[top])] + strconv.Itoa(current)
			}

		}
	}
	for i, size := 0, len(stack); i < size; i++ {
		if len(stack[i]) == 1 {
			table[stack[i]] += 1
		}
		if len(stack[i]) == 2 && stack[i][1] >= 'a' && stack[i][1] <= 'z' {
			table[stack[i]] += 1
		}
		if len(stack[i]) > 2 && stack[i][1] >= 'a' && stack[i][1] <= 'z' {
			count, _ := strconv.Atoi(stack[i][2:])
			table[stack[i][0:2]] += count
		}
		if len(stack[i]) > 1 && unicode.IsNumber(rune(stack[i][1])) {
			count, _ := strconv.Atoi(stack[i][1:])
			table[stack[i][0:1]] += count
		}
	}
	for k, v := range table {
		if !unicode.IsLetter(rune(k[0])) {
			continue
		}
		if v != 1 {
			format := k + strconv.Itoa(v)
			restack = append(restack, format)
		} else {
			restack = append(restack, k)
		}
	}
	sort.Strings(restack)
	return strings.Join(restack, "")
}
