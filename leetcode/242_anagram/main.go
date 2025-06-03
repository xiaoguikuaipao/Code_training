package main

func main() {

}

func isAnagram(s string, t string) bool {
	tab := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for _, c := range s {
		tab[c]++
	}
	for _, c := range t {
		if tab[c]-1 < 0 {
			return false
		}
		tab[c]--
	}
	return true
}
