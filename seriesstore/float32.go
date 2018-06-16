package seriesstore

import (
	"sync"
)

// Float32SStore is a store of float32 slices
// Implements the SeriesStore interface
// All getter and setter functions provide bound checks where applicable
// Embedded sync.Mutex to provide atomic operation ability
type Float32SStore struct {
	sync.Mutex
	store map[string][]float32
}

// NewFloat32SStore constructs and initializes a new Float32SStore
// Always use this function to init new Float32SStores
func NewFloat32SStore() *Float32SStore {
	return &Float32SStore{store: make(map[string][]float32)}
}

func (s *Float32SStore) set(key string, value []float32) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key in the store
func (s *Float32SStore) Set(key string, value []float32) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Float32SStore) setIdx(key string, idx int, value float32) error {
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
func (s *Float32SStore) SetIdx(key string, idx int, value float32) error {
	s.Lock()
	err := s.setIdx(key, idx, value)
	s.Unlock()

	return err
}

func (s *Float32SStore) get(key string) ([]float32, bool) {
	// explicitly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get returns the value for the given key
func (s *Float32SStore) Get(key string) ([]float32, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Float32SStore) getIdx(key string, idx int) (float32, error) {
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

// GetIdx returns the value for the given key at the specified index
func (s *Float32SStore) GetIdx(key string, idx int) (float32, error) {
	s.Lock()
	v, err := s.getIdx(key, idx)
	s.Unlock()

	return v, err
}

func (s *Float32SStore) getRange(key string, lower, upper int) ([]float32, error) {
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

// GetRange returns all values for the given key within the specified range (inclusive:exclusive)
func (s *Float32SStore) GetRange(key string, lower, upper int) ([]float32, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *Float32SStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *Float32SStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Float32SStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all keys of the store
func (s *Float32SStore) Members() []string {
	s.Lock()
	v := s.members()
	s.Unlock()

	return v
}

func (s *Float32SStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *Float32SStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Float32SStore) memberLen(key string) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	return len(v), nil
}

// MemberLen returns the length of the series value stored at the given key
func (s *Float32SStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *Float32SStore) clear() {
	s.store = make(map[string][]float32)
}

// Clear deletes all keys in the store
func (s *Float32SStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
