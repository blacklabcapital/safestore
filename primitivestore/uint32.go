package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of Uint32s
// Embedded sync.Mutex to provide atomic operation ability
type Uint32Store struct {
	sync.Mutex
	store map[string]uint32
}

func NewUint32Store() *Uint32Store {
	return &Uint32Store{store: make(map[string]uint32)}
}

func (s *Uint32Store) set(key string, value uint32) {
	s.store[key] = value
}

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

func (s *Uint32Store) Get(key string) (uint32, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Uint32Store) size() int {
	return len(s.store)
}

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

func (s *Uint32Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Uint32Store) clear() {
	s.store = make(map[string]uint32)
}

func (s *Uint32Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
