package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of booleans
// Embedded sync.Mutex to provide atomic operation ability
type BoolStore struct {
	sync.Mutex
	store map[string]bool
}

func NewBoolStore() *BoolStore {
	return &BoolStore{store: make(map[string]bool)}
}

func (s *BoolStore) set(key string, value bool) {
	s.store[key] = value
}

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

func (s *BoolStore) Get(key string) (bool, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *BoolStore) size() int {
	return len(s.store)
}

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

func (s *BoolStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *BoolStore) clear() {
	s.store = make(map[string]bool)
}

func (s *BoolStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
