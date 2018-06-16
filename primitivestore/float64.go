package primitivestore

import (
	"sync"
)

// Float64Store is a store of float64s
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type Float64Store struct {
	sync.Mutex
	store map[string]float64
}

// NewFloat64Store constructs and initializes a new Float64Store
// Always use this function when creating a new Float64Store
func NewFloat64Store() *Float64Store {
	return &Float64Store{store: make(map[string]float64)}
}

func (s *Float64Store) set(key string, value float64) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
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

// Get returns the value for the given key
func (s *Float64Store) Get(key string) (float64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Float64Store) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
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

// Members returns all keys of the store
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

// IsMember checks if the given key exists in the store
func (s *Float64Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Float64Store) clear() {
	s.store = make(map[string]float64)
}

// Clear deletes all keys in the store
func (s *Float64Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
