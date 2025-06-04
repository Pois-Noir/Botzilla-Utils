package Ren

import (
	"errors"
	"sync"
)

type SafeList[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]*V
}

func NewSafeList[K comparable, V any]() *SafeList[K, V] {
	return &SafeList[K, V]{
		data: make(map[K]*V),
	}
}

func (s *SafeList[K, V]) Get(key K) *V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := s.data[key]

	return result
}

func (s *SafeList[K, V]) Add(key K, value *V) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	search := s.data[key]
	if search != nil {
		return errors.New("key already exists")
	}

	s.data[key] = value

	return nil
}

func (s *SafeList[K, V]) Remove(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = nil
}
