package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of Ints
// Embedded sync.Mutex to provide atomic operation ability
type IntStore struct {
	sync.Mutex
	store map[string]int
}

func NewIntStore() *IntStore {
	return &IntStore{store: make(map[string]int)}
}

func (s *IntStore) set(key string, value int) {
	s.store[key] = value
}

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

func (s *IntStore) Get(key string) (int, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *IntStore) size() int {
	return len(s.store)
}

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

func (s *IntStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *IntStore) clear() {
	s.store = make(map[string]int)
}

func (s *IntStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
