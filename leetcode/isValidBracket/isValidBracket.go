package isValidBracket

func isValid(s string) bool {
	st := make([]byte, 0)
	for _, c := range s {
		switch c {
		case '[', '(', '{':
			st = append(st, byte(c))
		case ']', ')', '}':
			switch c {
			case ']':
				if len(st) == 0 || st[len(st)-1] != '[' {
					return false
				}
				st = st[:len(st)-1]
			case ')':
				if len(st) == 0 || st[len(st)-1] != '(' {
					return false
				}
				st = st[:len(st)-1]
			case '}':
				if len(st) == 0 || st[len(st)-1] != '{' {
					return false
				}
				st = st[:len(st)-1]
			}
		}
	}
	if len(st) > 0 {
		return false
	}
	return true
}
