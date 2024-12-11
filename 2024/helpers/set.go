package helpers

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{
		values: make(map[T]struct{}),
	}
	for _, e := range elements {
		set.Add(e)
	}
	return set
}

func (s *Set[T]) Add(value T) {
	s.values[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

func (s *Set[T]) Has(value T) bool {
	_, found := s.values[value]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Iter() func(func(item T) bool) {
	return func(yield func(item T) bool) {
		for k := range s.values {
			if !yield(k) {
				return
			}
		}
	}
}
