package queue

import (
	"container/heap"
)

// Interface .
type Interface interface {
	Less(any interface{}) bool
}

type sorter []Interface

func (s *sorter) Less(i, j int) bool {
	return (*s)[i].Less((*s)[j])
}
func (s *sorter) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func (s *sorter) Len() int {
	return len(*s)
}
func (s *sorter) Push(x interface{}) {
	*s = append((*s), x.(Interface))
}
func (s *sorter) Pop() (x interface{}) {
	n := s.Len()
	if n > 0 {
		x = (*s)[n-1]
		(*s) = (*s)[0 : n-1]
		return x
	}
	return nil
}

// PriorityQueue .
type PriorityQueue struct {
	s *sorter
}

// NewPriorityQueue .
func NewPriorityQueue() *PriorityQueue {
	q := new(PriorityQueue)
	q.s = new(sorter)
	heap.Init(q.s)
	return q
}

// Push .
func (q *PriorityQueue) Push(x Interface) {
	heap.Push(q.s, x)
}

// Pop .
func (q *PriorityQueue) Pop() Interface {
	return heap.Pop(q.s).(Interface)
}

// Top .
func (q *PriorityQueue) Top() Interface {
	if len(*q.s) > 0 {
		return (*q.s)[0].(Interface)
	}
	return nil
}

// Fix .
func (q *PriorityQueue) Fix(x Interface, i int) {
	(*q.s)[i] = x
	heap.Fix(q.s, i)
}

// Remove .
func (q *PriorityQueue) Remove(i int) Interface {
	return heap.Remove(q.s, i).(Interface)
}

// Len .
func (q *PriorityQueue) Len() int {
	return q.s.Len()
}

// GetSource .
func (q *PriorityQueue) GetSource() (ret []Interface) {
	if q.s != nil {
		for i := range *q.s {
			ret = append(ret, (*q.s)[i])
		}
		return
	}
	return nil
}
