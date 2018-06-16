package primitivestore

import (
	"sync"
)

// IntStore is a store of ints
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type IntStore struct {
	sync.Mutex
	store map[string]int
}

// NewIntStore constructs and initializes a new IntStore
// Always use this function when creating a new IntStore
func NewIntStore() *IntStore {
	return &IntStore{store: make(map[string]int)}
}

func (s *IntStore) set(key string, value int) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
func (s *IntStore) Set(key string, value int) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *IntStore) get(key string) (int, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *IntStore) Get(key string) (int, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *IntStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *IntStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *IntStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *IntStore) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *IntStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *IntStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *IntStore) clear() {
	s.store = make(map[string]int)
}

// Clear deletes all keys in the store
func (s *IntStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
