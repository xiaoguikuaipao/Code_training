package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	vhead := new(ListNode)
	vhead.Next = head
	start := vhead
	end := vhead
	for end != nil && end.Next != nil {
		end = end.Next.Next
		start = start.Next
		if end == start {
			goto FIND
		}
	}
	return nil

FIND:
	start = head
	for start != end {
		start = start.Next
		end = end.Next
	}
	return start
}
