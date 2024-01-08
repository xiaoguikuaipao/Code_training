package removeDuplicates

func removeDuplicates(s string) string {
	st := make([]byte, 0)
	for _, e := range s {
		st = append(st, byte(e))
		for len(st) > 1 && st[len(st)-1] == st[len(st)-2] {
			st = st[:len(st)-2]
		}
	}
	return string(st)
}
