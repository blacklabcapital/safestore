package primitivestore

import (
	"sync"
)

// BoolStore is a store of booleans
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type BoolStore struct {
	sync.Mutex
	store map[string]bool
}

// NewBoolStore constructs and initializes a new BoolStore
// Always use this function when creating a new BoolStore
func NewBoolStore() *BoolStore {
	return &BoolStore{store: make(map[string]bool)}
}

func (s *BoolStore) set(key string, value bool) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
func (s *BoolStore) Set(key string, value bool) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *BoolStore) get(key string) (bool, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *BoolStore) Get(key string) (bool, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *BoolStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *BoolStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *BoolStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *BoolStore) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *BoolStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *BoolStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *BoolStore) clear() {
	s.store = make(map[string]bool)
}

// Clear deletes all keys in the store
func (s *BoolStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
