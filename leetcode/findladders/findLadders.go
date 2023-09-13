package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

type state struct {
	currentWord string
	path        []string
	endWord     string
	wordList    []string
}

var res [][]string
var minPath map[string]string

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	quit := false
	minPath = make(map[string]string)
	res = make([][]string, 0)
	for _, target := range wordList {
		if endWord == target {
			quit = true
		}
	}
	if quit == false {
		return [][]string{}
	}
	queue := make([]state, 0)
	queue = append(queue, state{
		currentWord: beginWord,
		path:        []string{},
		endWord:     endWord,
		wordList:    wordList,
	})
	deal(queue)
	return res
}

func deal(queue []state) {
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentWord := current.currentWord
		path := make([]string, len(current.path))
		copy(path, current.path)
		path = append(path, currentWord)
		if len(res) > 0 && len(path) > len(res[0]) {
			return
		}
		if currentWord == current.endWord {
			res = append(res, path)
			pLen := len(path)
			for i := 0; i < pLen-1; i++ {
				minPath[path[i]] = path[i+1]
			}
		}
		wordList := current.wordList
		nextList := findNext(currentWord, wordList)
		if len(nextList) == 0 {
			continue
		}
		for _, nextWord := range nextList {
			repeated := false
			for _, already := range path {
				if already == nextWord {
					repeated = true
				}
			}
			if repeated == true {
				continue
			}
			if _, ok := minPath[nextWord]; ok {
				nexthop := minPath[nextWord]
				for {
					if nexthop == current.endWord {
						newState := state{
							currentWord: nexthop,
							path:        path,
							endWord:     current.endWord,
							wordList:    current.wordList,
						}
						queue = append(queue, newState)
						break
					} else {
						nexthop = minPath[nexthop]
						path = append(path, nexthop)
					}
				}
			} else {
				nextState := state{
					currentWord: nextWord,
					path:        path,
					endWord:     current.endWord,
					wordList:    current.wordList,
				}
				queue = append(queue, nextState)
			}
		}
	}
}

func findNext(currentWord string, wordList []string) (nextList []string) {
	cnt := 0
	n := len(currentWord)
	for _, Word := range wordList {
		cnt = 0
		for i := 0; i < n; i++ {
			if currentWord[i] != Word[i] {
				cnt++
			}
			if cnt > 1 {
				break
			}
		}
		if cnt == 1 {
			nextList = append(nextList, Word)
		}
	}
	return
}
