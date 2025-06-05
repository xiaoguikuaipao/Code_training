package datastructimpl

import (
	"sync"
)

type Dequer[T any] interface {
	FrontOp[T]
	BackOp[T]
	LenInfoer[T]
	Shower[T]
}

type FrontOp[T any] interface {
	PopFront() (T, bool)
	PushFront(T) bool
	Front() (T, bool)
}
type BackOp[T any] interface {
	PopBack() (T, bool)
	PushBack(T) bool
	Back() (T, bool)
}

type LenInfoer[T any] interface {
	Len() int
	Cap() int
	Empty() bool
	Full() bool
}

type Shower[T any] interface {
	Show() []T
}

type Deque[T any] struct {
	items []T
	cap   int
}

func (d *Deque[T]) PopBack() (T, bool) {
	if !d.Empty() {
		result := d.items[len(d.items)-1]
		d.items = d.items[:len(d.items)-1]
		return result, true
	}
	var zero T
	return zero, false
}

func (d *Deque[T]) PushBack(item T) bool {
	if d.items == nil {
		d.items = make([]T, 0, d.cap)
	}
	if !d.Full() {
		d.items = append(d.items, item)
		return true
	}
	return false
}

func (d *Deque[T]) Back() (T, bool) {
	if !d.Empty() {
		return d.items[len(d.items)-1], true
	}
	var zero T
	return zero, false
}

func (d *Deque[T]) PopFront() (T, bool) {
	if !d.Empty() {
		result := d.items[0]
		d.items = d.items[1:]
		return result, true
	}
	var zero T
	return zero, false
}

func (d *Deque[T]) PushFront(item T) bool {
	if d.items == nil {
		d.items = make([]T, 0, d.cap)
	}
	if !d.Full() {
		d.items = append([]T{item}, d.items...)

		return true
	}
	return false
}

func (d *Deque[T]) Front() (T, bool) {
	if !d.Empty() {
		return d.items[0], true
	}
	var zero T
	return zero, false
}

func (d *Deque[T]) Empty() bool {
	return len(d.items) == 0
}

func (d *Deque[T]) Full() bool {
	if d.cap == 0 {
		return false
	}
	if d.cap == len(d.items) {
		return true
	}
	return false
}

func (d *Deque[T]) Show() []T {
	return d.items
}

func (d *Deque[T]) Cap() int {
	return d.cap
}

func (d *Deque[T]) Len() int {
	return len(d.items)
}

type SyncDeque[T any] struct {
	Dequer[T]
	mu sync.RWMutex
}

func (d *SyncDeque[T]) PushBack(item T) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.Dequer.Full() {
		d.Dequer.PushBack(item)
		return true
	}
	return false
}

func (d *SyncDeque[T]) PopBack() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.Dequer.Empty() {
		return d.Dequer.PopBack()
	}
	var zero T
	return zero, false
}

func (d *SyncDeque[T]) Back() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.Dequer.Back()
}

func (d *SyncDeque[T]) PushFront(item T) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.Dequer.Full() {
		d.Dequer.PushFront(item)
		return true
	}
	return false
}

func (d *SyncDeque[T]) PopFront() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.Dequer.Empty() {
		return d.Dequer.PopFront()
	}
	var zero T
	return zero, false
}

func (d *SyncDeque[T]) Front() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.Dequer.Empty() {
		return d.Dequer.Front()
	}
	var zero T
	return zero, false
}

func NewDeque[T any](size ...int) Dequer[T] {
	d := &Deque[T]{}
	if len(size) > 0 && size[0] > 0 {
		d.cap = size[0]
	}
	d.items = make([]T, 0, d.cap)
	return d
}

func NewSyncDeque[T any](size ...int) Dequer[T] {
	d := &SyncDeque[T]{
		Dequer: NewDeque[T](size...),
	}
	return d
}

type Stack[T any] struct {
	d Dequer[T]
}

func NewStack[T any](size ...int) *Stack[T] {
	return &Stack[T]{
		d: NewDeque[T](size...),
	}
}

func NewSyncStack[T any](size ...int) *Stack[T] {
	return &Stack[T]{
		d: NewSyncDeque[T](size...),
	}
}

func (s *Stack[T]) Push(item T) bool {
	return s.d.PushBack(item)
}

func (s *Stack[T]) Peak() (T, bool) {
	return s.d.Back()
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.d.PopBack()
}

func (s *Stack[T]) Empty() bool {
	return s.d.Empty()
}

func (s *Stack[T]) Full() bool {
	return s.d.Full()
}

func (s *Stack[T]) Show() []T {
	return s.d.Show()
}

func (s *Stack[T]) Len() int {
	return s.d.Len()
}

type Queue[T any] struct {
	d Dequer[T]
}

func (q *Queue[T]) Enqueue(item T) bool {
	return q.d.PushBack(item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.d.PopFront()
}
func (q *Queue[T]) Peak() (T, bool) {
	return q.d.Front()
}
func (q *Queue[T]) Empty() bool {
	return q.d.Empty()
}
func (q *Queue[T]) Full() bool {
	return q.d.Full()
}
func (q *Queue[T]) Show() []T {
	return q.d.Show()
}

func (q *Queue[T]) Len() int {
	return q.d.Len()
}
