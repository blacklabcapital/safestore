package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of Int64s
// Embedded sync.Mutex to provide atomic operation ability
type Int64Store struct {
	sync.Mutex
	store map[string]int64
}

func NewInt64Store() *Int64Store {
	return &Int64Store{store: make(map[string]int64)}
}

func (s *Int64Store) set(key string, value int64) {
	s.store[key] = value
}

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

func (s *Int64Store) Get(key string) (int64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Int64Store) size() int {
	return len(s.store)
}

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

func (s *Int64Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Int64Store) clear() {
	s.store = make(map[string]int64)
}

func (s *Int64Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
