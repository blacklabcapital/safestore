package seriesstore

import (
	"sync"
)

// IntSStore is a store of int slices
// Implements the SeriesStore interface
// All getter and setter functions provide bound checks where applicable
// Embedded sync.Mutex to provide atomic operation ability
type IntSStore struct {
	sync.Mutex
	store map[string][]int
}

// NewIntSStore constructs and initializes a new IntSStore
// Always use this function to init new IntSStores
func NewIntSStore() *IntSStore {
	return &IntSStore{store: make(map[string][]int)}
}

func (s *IntSStore) set(key string, value []int) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key in the store
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

// SetIdx stores the given value mapped to the given key at the specified index in the store
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

// Get accesses the value for the given key
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

// GetIdx accesses the value for the given key at the specified index
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

// GetRange gets all values for the given key within the specified range (inclusive:exclusive)
func (s *IntSStore) GetRange(key string, lower, upper int) ([]int, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *IntSStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
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

// Members returns all member keys of the store
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

// IsMember checks if the given key exists in the store
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

// MemberLen returns the length of the series value stored at the given key
func (s *IntSStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *IntSStore) clear() {
	s.store = make(map[string][]int)
}

// Clear deletes all keys in the store
func (s *IntSStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
