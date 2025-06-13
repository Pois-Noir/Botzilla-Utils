package safemap

import (
	"errors"
	"maps"
	"sync"
)

type SafeMap[K comparable, V any] struct { // file name changed list-> map
	mu   sync.RWMutex
	data map[K]*V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]*V),
	}
}

func (s *SafeMap[K, V]) Get(key K) *V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := s.data[key]

	return result
}

func (s *SafeMap[K, V]) Add(key K, value *V) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	search := s.data[key]
	if search != nil {
		return errors.New("key already exists")
	}

	s.data[key] = value

	return nil
}

func (s *SafeMap[K, V]) Remove(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = nil
}

func (s *SafeMap[K, V]) Data() map[K]*V {
	s.mu.RLock()
	defer s.mu.RUnlock()

	copy := make(map[K]*V, len(s.data))

	maps.Copy(copy, s.data)

	return copy
}

func (s *SafeMap[K, V]) ForEach(f func(K, *V)) {

	// Copy of original data
	data := s.Data()

	for key, val := range data {
		f(key, val)
	}
}
