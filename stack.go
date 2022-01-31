// Package stack provides simple, thread-safe stack.
//
// Stacks are compose of any primitive types (bool, int, float32, string, etc), or any data types.
package stack

// Stack is the common interface for all stack implementation.
type Stack[V any] interface {
	Push(value V)
	Pop() (value V, ok bool)
	Peek() (value V, ok bool)

	Size() int
}

// New returns new thread-safe stack.
func New[V any]() Stack[V] {
	return newSafeStack(NewUnsafe[V]())
}

func NewUnsafe[V any]() Stack[V] {
	return newArrayStack[V]()
}
