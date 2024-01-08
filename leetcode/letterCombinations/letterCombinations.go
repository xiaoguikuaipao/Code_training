package letterCombinations

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	current := make([]byte, 0)
	if len(digits) == 0 {
		return []string{}
	}
	table := make(map[byte][]byte)
	table['2'] = []byte{'a', 'b', 'c'}
	table['3'] = []byte{'d', 'e', 'f'}
	table['4'] = []byte{'g', 'h', 'i'}
	table['5'] = []byte{'j', 'k', 'l'}
	table['6'] = []byte{'m', 'n', 'o'}
	table['7'] = []byte{'p', 'q', 'r', 's'}
	table['8'] = []byte{'t', 'u', 'v'}
	table['9'] = []byte{'w', 'x', 'y', 'z'}
	for _, e := range table[digits[0]] {
		current = append(current, e)
		backtracking(0, current, &ret, digits, &table)
		current = current[:len(current)-1]
	}
	return ret
}

func backtracking(i int, current []byte, ret *[]string, digits string, table *map[byte][]byte) {
	i++
	if i > len(digits)-1 {
		*ret = append(*ret, string(current))
		return
	}

	for _, e := range (*table)[digits[i]] {
		current = append(current, e)
		backtracking(i, current, ret, digits, table)
		current = current[:len(current)-1]
	}

}
