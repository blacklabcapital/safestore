package seriesstore

import (
	"sync"
)

// Implements the SeriesStore interface
// A store of int slices
// All get and set functions provide bound checks
// Embedded sync.Mutex to provide atomic operation ability
type IntSStore struct {
	sync.Mutex
	store map[string][]int
}

func NewIntSStore() *IntSStore {
	return &IntSStore{store: make(map[string][]int)}
}

func (s *IntSStore) set(key string, value []int) {
	s.store[key] = value
}

func (s *IntSStore) Set(key string, value []int) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *IntSStore) setIdx(key string, idx int, value int) error {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return ErrKeyDoesNotExist
	}

	// bounds check
	if idx < 0 || idx >= len(v) {
		return ErrIdxOutOfBounds
	}

	s.store[key][idx] = value

	return nil
}

func (s *IntSStore) SetIdx(key string, idx int, value int) error {
	s.Lock()
	err := s.setIdx(key, idx, value)
	s.Unlock()

	return err
}

func (s *IntSStore) get(key string) ([]int, bool) {
	// explicitly return second return value
	v, ok := s.store[key]

	return v, ok
}

func (s *IntSStore) Get(key string) ([]int, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *IntSStore) getIdx(key string, idx int) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0.0, ErrKeyDoesNotExist
	}

	// bounds check
	if idx < 0 || idx >= len(v) {
		return 0.0, ErrIdxOutOfBounds
	}

	return s.store[key][idx], nil
}

func (s *IntSStore) GetIdx(key string, idx int) (int, error) {
	s.Lock()
	v, err := s.getIdx(key, idx)
	s.Unlock()

	return v, err
}

func (s *IntSStore) getRange(key string, lower, upper int) ([]int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return nil, ErrKeyDoesNotExist
	}

	// bounds check
	if lower < 0 || lower > len(v) || upper < 0 || upper > len(v) {
		return nil, ErrIdxOutOfBounds
	}

	return s.store[key][lower:upper], nil
}

func (s *IntSStore) GetRange(key string, lower, upper int) ([]int, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *IntSStore) size() int {
	return len(s.store)
}

func (s *IntSStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *IntSStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

func (s *IntSStore) Members() []string {
	s.Lock()
	v := s.members()
	s.Unlock()

	return v
}

func (s *IntSStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

func (s *IntSStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *IntSStore) memberLen(key string) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	return len(v), nil
}

func (s *IntSStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *IntSStore) clear() {
	s.store = make(map[string][]int)
}

func (s *IntSStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
