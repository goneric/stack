package stack

import "testing"

func TestStack_Len(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	size := s.Size()
	if size != 3 {
		t.Errorf("wrong size, expect %d, got %d", 3, size)
	}
}

func TestStack_Peek(t *testing.T) {
	s := New[string]()

	// Peek empty stack
	_, ok := s.Peek()
	if ok {
		t.Errorf("expect ok is false for empty stack")
	}

	s.Push("task 1")
	s.Push("task 2")
	s.Push("task 3")
	task, ok := s.Peek()
	if !ok || task != "task 3" {
		t.Errorf("expect task 3, got '%s'", task)
	}
	size := s.Size()
	if size != 3 {
		t.Errorf("wrong size, expect %d, got %d", 3, size)
	}
}

func TestStack_Pop(t *testing.T) {
	s := New[string]()

	// Pop empty stack
	_, ok := s.Pop()
	if ok {
		t.Errorf("expect ok is false for empty stack")
	}

	s.Push("task 1")
	s.Push("task 2")
	s.Push("task 3")
	task, ok := s.Pop()
	if !ok || task != "task 3" {
		t.Errorf("expect task 3, got '%s'", task)
	}
	size := s.Size()
	if size != 2 {
		t.Errorf("wrong size, expect %d, got %d", 2, size)
	}
}
