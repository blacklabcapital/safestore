package primitivestore

import (
	"sync"
)

// Int64Store is a store of Int64s
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type Int64Store struct {
	sync.Mutex
	store map[string]int64
}

// NewInt64Store constructs and initializes a new Int64Store
// Always use this function when creating a new Int64Store
func NewInt64Store() *Int64Store {
	return &Int64Store{store: make(map[string]int64)}
}

func (s *Int64Store) set(key string, value int64) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
func (s *Int64Store) Set(key string, value int64) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Int64Store) get(key string) (int64, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *Int64Store) Get(key string) (int64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Int64Store) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *Int64Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Int64Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *Int64Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Int64Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *Int64Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Int64Store) clear() {
	s.store = make(map[string]int64)
}

// Clear deletes all keys in the store
func (s *Int64Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
