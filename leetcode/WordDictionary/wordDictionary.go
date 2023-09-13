package main

import "fmt"

func main() {
	wd := Constructor()
	wd.AddWord("hello")
	fmt.Println(wd.Search("hello"))
	wd.AddWord("hallo")
	fmt.Println(wd.Search("hallo"))
	wd.AddWord("a")
	fmt.Println(wd.Search("a"))
}

type WordDictionary struct {
	children map[rune]*WordDictionary
	isWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children: make(map[rune]*WordDictionary),
		isWord:   false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	parent := this
	for _, literal := range word {
		if _, ok := parent.children[literal]; !ok {
			parent.children[literal] = &WordDictionary{
				children: make(map[rune]*WordDictionary),
				isWord:   false,
			}
		}
		parent = parent.children[literal]
	}
	// 这里是单词的标志(结尾)
	parent.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	parent := this
	for i, literal := range word {
		if rune(literal) == '.' {
			// 先判断为false，如果search后面返回true，就是true
			isMatched := false
			for _, v := range parent.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if _, ok := parent.children[literal]; !ok {
			return false
		}
		parent = parent.children[literal]
	}
	return len(parent.children) == 0 || parent.isWord
}
