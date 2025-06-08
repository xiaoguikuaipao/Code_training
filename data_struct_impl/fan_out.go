package datastructimpl

import (
	"context"
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
)

type ITaskPool[T Itask] interface {
	AddRobin(context.Context, T) error
	AddHash(context.Context, T) error
}

type TaskPool[T Itask] []chan T

func (p TaskPool[T]) AddRobin(ctx context.Context, t T) error {
	select {
	case p[int(t.ID())%len(p)] <- t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p TaskPool[T]) AddHash(ctx context.Context, t T) error {
	hash := crc32.ChecksumIEEE([]byte(t.Attribute4Hash()))
	index := hash % uint32(len(p))
	select {
	case p[index] <- t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Itask every task is supposed to have a deal method, implement it
type Itask interface {
	ID() uint32
	Attribute4Hash() string
	Deal() error
}

type BaseTask struct {
}

func (b *BaseTask) ID() uint64 {
	return uint64(rand.Int())
}

func (b *BaseTask) Attribute4Hash() string {
	return strconv.Itoa(rand.Int())
}

func (b *BaseTask) Deal() error {
	return nil
}

// Iworker every worker need to implement work()
type Iworker interface {
	work()
}

// defaultWorker only has a ID, and simply deal task from task queue, dealing with task's deal()
type defaultWorker[T Itask] struct {
	queueID int
	ID      int
}

func (w *defaultWorker[T]) work(queue chan T) {
	fmt.Printf("queue[%d].worker[%d] started...\n", w.queueID, w.ID)
	for t := range queue {
		err := t.Deal()
		if err != nil {
			fmt.Printf("queue[%d].worker[%d] Deal task error\n", w.queueID, w.ID)
		} else {
			fmt.Printf("queue[%d].worker[%d] Deal task success\n", w.queueID, w.ID)
		}
	}
}

// NewQueueGroup represent a general fan-out concurrent model
func NewQueueGroup[T Itask](maxQueueCnt, maxTaskInQ, maxWorkerCnt int) TaskPool[T] {
	qg := make([]chan T, maxQueueCnt)
	for i := range qg {
		qg[i] = make(chan T, maxTaskInQ)
		for j := range maxWorkerCnt {
			worker := &defaultWorker[T]{
				queueID: i,
				ID:      j,
			}
			go worker.work(qg[i])
		}
	}
	return qg
}
