package stack

type arrayStack[V any] struct {
	values []V
}

func newArrayStack[V any]() *arrayStack[V] {
	return &arrayStack[V]{
		values: make([]V, 0),
	}
}

func (s *arrayStack[V]) Push(value V) {
	s.values = append(s.values, value)
}

func (s *arrayStack[V]) Pop() (value V, ok bool) {
	if len(s.values) == 0 {
		ok = false
		return
	}
	value = s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return value, true
}

func (s *arrayStack[V]) Peek() (value V, ok bool) {
	if len(s.values) == 0 {
		ok = false
		return
	}
	value = s.values[len(s.values)-1]
	return value, true
}

func (s *arrayStack[V]) Size() int {
	return len(s.values)
}
