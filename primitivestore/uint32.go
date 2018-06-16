package primitivestore

import (
	"sync"
)

// Uint32Store is a store of Uint32s
// Implements the PrimitiveStore interface
// Embedded sync.Mutex to provide atomic operation ability
type Uint32Store struct {
	sync.Mutex
	store map[string]uint32
}

// NewUint32Store constructs and initializes a new Uint32Store
func NewUint32Store() *Uint32Store {
	return &Uint32Store{store: make(map[string]uint32)}
}

func (s *Uint32Store) set(key string, value uint32) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key
func (s *Uint32Store) Set(key string, value uint32) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Uint32Store) get(key string) (uint32, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *Uint32Store) Get(key string) (uint32, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Uint32Store) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *Uint32Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Uint32Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *Uint32Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Uint32Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *Uint32Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Uint32Store) clear() {
	s.store = make(map[string]uint32)
}

// Clear deletes all keys in the store
func (s *Uint32Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
