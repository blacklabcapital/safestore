package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of Int32s
// Embedded sync.Mutex to provide atomic operation ability
type Int32Store struct {
	sync.Mutex
	store map[string]int32
}

func NewInt32Store() *Int32Store {
	return &Int32Store{store: make(map[string]int32)}
}

func (s *Int32Store) set(key string, value int32) {
	s.store[key] = value
}

func (s *Int32Store) Set(key string, value int32) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Int32Store) get(key string) (int32, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

func (s *Int32Store) Get(key string) (int32, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Int32Store) size() int {
	return len(s.store)
}

func (s *Int32Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Int32Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

func (s *Int32Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Int32Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

func (s *Int32Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Int32Store) clear() {
	s.store = make(map[string]int32)
}

func (s *Int32Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
