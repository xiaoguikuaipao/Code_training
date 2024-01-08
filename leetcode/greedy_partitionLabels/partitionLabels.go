package greedy_partitionLabels

type be struct {
	Begin, End int
}

func partitionLabels(s string) []int {
	table := make(map[byte]be)
	ret := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := table[s[i]]; ok {
			beV := table[s[i]]
			beV.End = i
			table[s[i]] = beV
		} else {
			beV := be{Begin: i, End: i}
			table[s[i]] = beV
		}
	}
	start := table[s[0]].Begin
	end := table[s[0]].End
	for i := 1; i < len(s); i++ {
		if table[s[i]].Begin < end {
			if table[s[i]].End > end {
				end = table[s[i]].End
			}
		} else {
			ret = append(ret, end-start+1)
			start = table[s[i]].Begin
			end = table[s[i]].End
		}
	}
	ret = append(ret, end-start+1)
	return ret
}
