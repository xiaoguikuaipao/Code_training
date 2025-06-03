package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	vhead := new(ListNode)
	vhead.Next = head
	prev := vhead
	curr := head
	for curr != nil {
		succ := curr.Next
		curr.Next = prev
		prev = curr
		curr = succ
	}
	head.Next = nil
	vhead.Next = prev
	return vhead.Next
}
