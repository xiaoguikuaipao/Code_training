package main

import "fmt"

func main() {
	ll := Constructor()
	ll.Print()
	ll.AddAtHead(1)
	ll.Print()
	ll.AddAtTail(3)
	ll.Print()
	ll.AddAtIndex(1, 2)
	ll.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	rhead *ListNode
	head  *ListNode
	num   int
}

func Constructor() MyLinkedList {
	ret := MyLinkedList{
		num: 0,
	}
	ret.rhead = &ListNode{Next: ret.head}

	return ret
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.num {
		return -1
	}
	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{Val: val, Next: this.head}
	this.rhead.Next = newHead
	this.head = newHead
	this.num++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newTail := &ListNode{Val: val}
	if this.head == nil {
		this.head = newTail
		this.rhead = newTail
		this.num++
	}
	curr := this.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newTail
	this.num++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.num {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.num {
		this.AddAtTail(val)
		return
	}
	curr := this.head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	prev := curr
	succ := curr.Next
	newNode := &ListNode{Val: val, Next: succ}
	prev.Next = newNode
	this.num++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.num {
		return
	}
	if index == 0 {
		newHead := this.head.Next
		this.rhead.Next = newHead
		this.head = newHead
	} else {
		curr := this.head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		prev := curr
		succ := curr.Next.Next
		prev.Next = succ
	}
	this.num--
}

func (this *MyLinkedList) Print() {
	curr := this.head
	for curr != nil {
		fmt.Printf("%d->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
