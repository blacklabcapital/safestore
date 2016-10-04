package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of float64s
// Embedded sync.Mutex to provide atomic operation ability
type Float64Store struct {
	sync.Mutex
	store map[string]float64
}

func NewFloat64Store() *Float64Store {
	return &Float64Store{store: make(map[string]float64)}
}

func (s *Float64Store) set(key string, value float64) {
	s.store[key] = value
}

func (s *Float64Store) Set(key string, value float64) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Float64Store) get(key string) (float64, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

func (s *Float64Store) Get(key string) (float64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Float64Store) size() int {
	return len(s.store)
}

func (s *Float64Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Float64Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

func (s *Float64Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Float64Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

func (s *Float64Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Float64Store) clear() {
	s.store = make(map[string]float64)
}

func (s *Float64Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
