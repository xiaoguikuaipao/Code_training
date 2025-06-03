package main

type House struct {
}

func enter(i int) int {
	print(i)
	return i
}
func main() {
	switch 1 {
	case enter(2), enter(3), enter(1), enter(4):
		fallthrough
	case enter(999):
		print("haha")
	}
}
