package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	vhead := new(ListNode)
	vhead.Next = head
	start := vhead
	end := vhead
	for i := 0; i < n && end != nil; i++ {
		end = end.Next
	}
	for end.Next != nil {
		start = start.Next
		end = end.Next
	}
	del := start.Next
	succ := del.Next
	start.Next = succ
	del.Next = nil
	return vhead.Next
}
