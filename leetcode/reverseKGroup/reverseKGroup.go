package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	Rhead := &ListNode{}
	Rhead.Next = head
	resEnd := &ListNode{}
	head, resEnd = reverseK(Rhead, k)
	for resEnd != nil {
		_, resEnd = reverseK(resEnd, k)
	}
	return head
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode) {
	resEnd := head.Next
	end := head.Next
	for i := 0; i < k-1; i++ {
		if end.Next == nil {
			return head.Next, nil
		}
		end = end.Next
	}
	resEnd.Next = end
	next := end.Next
	second := head.Next.Next
	head.Next.Next = next
	for i := 0; i < k-1; i++ {
		temp := second.Next
		second.Next = head.Next
		head.Next = second
		second = temp
	}
	return head.Next, resEnd
}
