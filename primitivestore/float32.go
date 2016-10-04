package primitivestore

import (
	"sync"
)

// Implements the PrimitiveStore interface
// A store of float32s
// Embedded sync.Mutex to provide atomic operation ability
type Float32Store struct {
	sync.Mutex
	store map[string]float32
}

func NewFloat32Store() *Float32Store {
	return &Float32Store{store: make(map[string]float32)}
}

func (s *Float32Store) set(key string, value float32) {
	s.store[key] = value
}

func (s *Float32Store) Set(key string, value float32) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Float32Store) get(key string) (float32, bool) {
	// explictly return second return value
	v, ok := s.store[key]

	return v, ok
}

func (s *Float32Store) Get(key string) (float32, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Float32Store) size() int {
	return len(s.store)
}

func (s *Float32Store) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Float32Store) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

func (s *Float32Store) Members() []string {
	s.Lock()
	mems := s.members()
	s.Unlock()

	return mems
}

func (s *Float32Store) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

func (s *Float32Store) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Float32Store) clear() {
	s.store = make(map[string]float32)
}

func (s *Float32Store) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
