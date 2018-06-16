package seriesstore

import (
	"sync"
)

// Float64SStore is a store of float64 slices
// Implements the SeriesStore interface
// All getter and setter functions provide bound checks where applicable
// Embedded sync.Mutex to provide atomic operation ability
type Float64SStore struct {
	sync.Mutex
	store map[string][]float64
}

// NewFloat64SStore constructs and initializes a new Float64SStore
// Always use this function to init new Float64SStores
func NewFloat64SStore() *Float64SStore {
	return &Float64SStore{store: make(map[string][]float64)}
}

func (s *Float64SStore) set(key string, value []float64) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key in the store
func (s *Float64SStore) Set(key string, value []float64) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Float64SStore) setIdx(key string, idx int, value float64) error {
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
func (s *Float64SStore) SetIdx(key string, idx int, value float64) error {
	s.Lock()
	err := s.setIdx(key, idx, value)
	s.Unlock()

	return err
}

func (s *Float64SStore) get(key string) ([]float64, bool) {
	// explicitly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get accesses the value for the given key
func (s *Float64SStore) Get(key string) ([]float64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Float64SStore) getIdx(key string, idx int) (float64, error) {
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
func (s *Float64SStore) GetIdx(key string, idx int) (float64, error) {
	s.Lock()
	v, err := s.getIdx(key, idx)
	s.Unlock()

	return v, err
}

func (s *Float64SStore) getRange(key string, lower, upper int) ([]float64, error) {
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
func (s *Float64SStore) GetRange(key string, lower, upper int) ([]float64, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *Float64SStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *Float64SStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Float64SStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all member keys of the store
func (s *Float64SStore) Members() []string {
	s.Lock()
	v := s.members()
	s.Unlock()

	return v
}

func (s *Float64SStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *Float64SStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Float64SStore) memberLen(key string) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	return len(v), nil
}

// MemberLen returns the length of the series value stored at the given key
func (s *Float64SStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *Float64SStore) clear() {
	s.store = make(map[string][]float64)
}

// Clear deletes all keys in the store
func (s *Float64SStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
