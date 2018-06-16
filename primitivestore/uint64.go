package primitivestore

import (
	"sync"
)

// Uint64Store is a store of Uint64s
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type Uint64Store struct {
	sync.Mutex
	store map[string]uint64
}

// NewUint64Store constructs and initializes a new Uint64Store
// Always use this function when creating a new Uint64Store
func NewUint64Store() *Uint64Store {
	return &Uint64Store{store: make(map[string]uint64)}
}

func (s *Uint64Store) set(key string, value uint64) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
func (s *Uint64Store) Set(key string, value uint64) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Uint64Store) get(key string) (uint64, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *Uint64Store) Get(key string) (uint64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Uint64Store) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *Uint64Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Uint64Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *Uint64Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Uint64Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *Uint64Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Uint64Store) clear() {
	s.store = make(map[string]uint64)
}

// Clear deletes all keys in the store
func (s *Uint64Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
