package state

import "sync"

type MemStorage[T comparable] struct {
	mu     sync.Mutex
	states map[int64]T
}

func NewMemStorage[T comparable]() *MemStorage[T] {
	return &MemStorage[T]{states: map[int64]T{}}
}

func (s *MemStorage[T]) Current(id int64) (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	state := s.states[id]

	return state, nil
}

func (s *MemStorage[T]) Set(id int64, name T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.states[id] = name

	return nil
}

func (s *MemStorage[T]) Clear(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.states, id)

	return nil
}
