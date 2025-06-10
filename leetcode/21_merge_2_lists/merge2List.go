package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := &ListNode{}
	head := cur
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}
	remain := list1
	if list1 == nil {
		remain = list2
	}
	cur.Next = remain
	return head.Next
}

func main() {

}
