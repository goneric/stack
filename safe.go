package stack

import (
	"sync"
)

type safeStack[V any] struct {
	lock *sync.RWMutex
	s    Stack[V]
}

func newSafeStack[V any](s Stack[V]) *safeStack[V] {
	return &safeStack[V]{s: s, lock: &sync.RWMutex{}}
}

func (s *safeStack[V]) Push(value V) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s.Push(value)
}

func (s *safeStack[V]) Pop() (value V, ok bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.s.Pop()
}

func (s *safeStack[V]) Peek() (value V, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.s.Peek()
}

func (s *safeStack[V]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.s.Size()
}
