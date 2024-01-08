package list_stack

type MyStack struct {
	list1 []int
	list2 []int
}

func Constructor() MyStack {
	return MyStack{
		list1: make([]int, 0),
		list2: make([]int, 0),
	}
}

func (this *MyStack) Top() int {
	if len(this.list1) == 0 {
		for len(this.list2) != 1 {
			this.list1 = append(this.list1, this.list2[0])
			this.list2 = this.list2[1:]
		}
		defer func() {
			this.list1 = append(this.list1, this.list2[0])
			this.list2 = this.list2[1:]
		}()
		return this.list2[0]
	}
	if len(this.list2) == 0 {
		for len(this.list1) != 1 {
			this.list2 = append(this.list2, this.list1[0])
			this.list1 = this.list1[1:]
		}
		defer func() {
			this.list2 = append(this.list2, this.list1[0])
			this.list1 = this.list1[1:]
		}()
		return this.list1[0]
	}
	return -1

}

func (this *MyStack) Push(x int) int {
	if len(this.list1) == 0 {
		this.list2 = append(this.list2, x)
	}
	if len(this.list2) == 0 {
		this.list1 = append(this.list1, x)
	}
	return 0
}

func (this *MyStack) Pop() int {
	if len(this.list1) == 0 {
		for len(this.list2) != 1 {
			this.list1 = append(this.list1, this.list2[0])
			this.list2 = this.list2[1:]
		}
		defer func() {
			this.list2 = this.list2[1:]
		}()
		return this.list2[0]
	}
	if len(this.list2) == 0 {
		for len(this.list1) != 1 {
			this.list2 = append(this.list2, this.list1[0])
			this.list1 = this.list1[1:]
		}
		defer func() {
			this.list1 = this.list1[1:]
		}()
		return this.list1[0]
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return len(this.list2) == 0 && len(this.list1) == 0
}
