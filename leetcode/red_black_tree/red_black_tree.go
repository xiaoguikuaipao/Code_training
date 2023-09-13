package main

import . "golang.org/x/exp/constraints"

func main() {

}

const (
	RED   = 0
	BLACK = 1
)

type node[K Ordered, V any] struct {
	left, right, parent *node[K, V]
	color               int
	Key                 K
	Value               V
}

type Tree[K Ordered, V any] struct {
	root *node[K, V]
	size int
}

func NewTree[K Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

func (t *Tree[K, V]) Find(key K) V {
	n := t.findnode(key)
	if n != nil {
		return n.Value
	}
	var result V
	return result
}

func (t *Tree[K, V]) FindIt(key K) *node[K, V] {
	return t.findnode(key)
}

func (t *Tree[K, V]) Size() int {
	return t.size
}

func (t *Tree[K, V]) Empty() bool {
	if t.root != nil {
		return true
	}
	return false
}

func (t *Tree[K, V]) Iterator() *node[K, V] {
	return minimum(t.root)
}

func (t *Tree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree[K, V]) Insert(key K, val V) {
	x := t.root
	var y *node[K, V]

	// 二分搜索, 结束后x为插入的叶子节点, y为该叶子的父节点
	for x != nil {
		y = x
		if key < x.Key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &node[K, V]{
		parent: y,
		Key:    key,
		Value:  val,
		color:  RED,
	}
	t.size++

	// y==nil 说明树是空的
	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if z.Key < y.Key {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)

}

func (t *Tree[K, V]) Delete(key K) {
	z := t.findnode(key)
	if z == nil {
		return
	}
	var x, y *node[K, V]
	//三种情况，z节点有两个子节点，有一个子节点，没有子节点
	// 如果找到的节点有左右孩子，不能直接删
	//改造的起点：y，要么是key节点本身，要么是比他大的第一个节点(右子树最左)
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	//如果上一步判断后，y是z，那么x就是z的唯一孩子
	//如果上一步判断后，y是比他大的第一个节点(右子树最左)，那么x可能是nil，可能是y的右孩子
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	// x!=nil的情况只有两种，一是x是z的唯一孩子，二是x是右子树最左的右孩子
	if x != nil {
		x.parent = xparent
	}
	//y的父亲是nil说明y是z，并且z是根, 那么x必然不是nil，是z的唯一一个孩子
	if y.parent == nil {
		t.root = x
		//y==y的父亲的左孩子，说明y是z，z是左孩子，或者y是z的右子树最左
	} else if y == y.parent.left {
		y.parent.left = x
		//y==y的父亲的右孩子，说明y是z，z是右孩子
	} else {
		y.parent.right = x
	}

	if y != z {
		z.Key = y.Key
		z.Value = y.Value
	}

	//红色节点可以直接删除，黑色节点需要fixup
	if y.color == BLACK {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

//两种情况：1. 删除的黑色节点的兄弟节点有红色子节点(用兄弟，左红，右红，左右红)
// 2. 删除的黑色节点的兄弟节点没有红色子节点(用父亲，父亲是黑->把黑父直接fixup->黑父的父是红,相当于第二种情况，父亲是红->把父亲变黑,兄弟变红)
//3. 删除的黑色节点有1个红色子节点
func (t *Tree[K, V]) rbDeleteFixup(x, xparent *node[K, V]) {
	var w *node[K, V]
	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			xparent = x.parent
		}
		if x == x.parent.left {
			w = x.parent.right
			if w.color == RED {
				w.color = BLACK
				xparent.color = RED
				t.leftRotate(xparent)
			}
		}
	}
}

func (t *Tree[K, V]) rbInsertFixup(x *node[K, V]) {

}

//对x节点左旋就是把x的右的左变成x的右，然后x挂在原本x的右的左孩子上
func (t *Tree[K, V]) leftRotate(x *node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

//对x节点右旋就是把x的左的右变成x的左，然后x挂在原本x的左的右孩子上
func (t *Tree[K, V]) rightRotate(x *node[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree[K, V]) findnode(key K) *node[K, V] {
	x := t.root
	for x != nil {
		if key < x.Key {
			x = x.left
		} else {
			if x.Key == key {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// 找到比给定节点大的最小节点
func successor[K Ordered, V any](x *node[K, V]) *node[K, V] {
	//要么是右子树的最小节点
	//要么是父亲节点
	//要么是nil
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

func minimum[K Ordered, V any](n *node[K, V]) *node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func getColor[K Ordered, V any](n *node[K, V]) int {
	if n != nil {
		return n.color
	}
	return BLACK
}
