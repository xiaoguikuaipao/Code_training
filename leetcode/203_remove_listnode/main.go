package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	hd := new(ListNode)
	hd.Next = head
	index := head
	prev := hd
	for index != nil {
		succ := index.Next
		if index.Val == val {
			prev.Next = succ
		} else {
			prev = prev.Next
		}
		index = succ
	}
	return hd.Next
}
