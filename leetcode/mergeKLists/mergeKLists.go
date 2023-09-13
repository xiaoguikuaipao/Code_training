package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//分治：先数组分治，再链表递归
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	num := len(lists) / 2
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	}
	right.Next = mergeTwoLists(left, right.Next)
	return right
}

// 暴力
//func mergeKLists(lists []*ListNode) *ListNode {
//	head := &ListNode{}
//	current := &ListNode{}
//	for {
//		count := 0
//		min := 1000
//		index := -1
//		for i := range lists {
//			if lists[i] != nil {
//				if lists[i].Val < min {
//					index = i
//					min = lists[i].Val
//				}
//				count++
//			}
//		}
//		if index != -1 {
//			if head.Next == nil {
//				current = lists[index]
//				head.Next = current
//			} else {
//				current.Next = lists[index]
//				current = current.Next
//			}
//			lists[index] = lists[index].Next
//		}
//		if count == 1 || index == -1 {
//			break
//		}
//	}
//	return head.Next
//}
