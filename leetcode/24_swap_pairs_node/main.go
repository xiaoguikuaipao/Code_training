package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
func swapPairs(head *ListNode) *ListNode {
	vhead := new(ListNode)
	vhead.Next = head
	if head == nil {
		return nil
	}
	prev := vhead
	curr := head
	succ := head.Next
	for succ != nil {
		prev.Next = succ
		curr.Next = succ.Next
		succ.Next = curr
		prev = curr
		curr = curr.Next
		if curr == nil {
			break
		}
		succ = curr.Next
	}
	return vhead.Next
}
