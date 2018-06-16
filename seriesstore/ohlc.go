package seriesstore

import (
	"sync"
)

type OHLC struct {
	Open  float32
	High  float32
	Low   float32
	Close float32
}

// OHLCSStore is a store of OHLC (Open High Low Close stock ticker prices) slices
// Implements the SeriesStore interface
// All getter and setter functions provide bound checks where applicable
// Embedded sync.Mutex to provide atomic operation ability
type OHLCSStore struct {
	sync.Mutex
	store map[string][]OHLC
}

// NewOHLCSStore constructs and initializes a new OHLCSStore
// Always use this function to init new OHLCSStore
func NewOHLCSStore() *OHLCSStore {
	return &OHLCSStore{store: make(map[string][]OHLC)}
}

func (s *OHLCSStore) set(key string, value []OHLC) {
	s.store[key] = value
}

// Set stores the given value mapped to the given key in the store
func (s *OHLCSStore) Set(key string, value []OHLC) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *OHLCSStore) setIdx(key string, idx int, value *OHLC) error {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return ErrKeyDoesNotExist
	}

	// bounds check
	if idx < 0 || idx >= len(v) {
		return ErrIdxOutOfBounds
	}

	s.store[key][idx] = *value

	return nil
}

// SetIdx stores the given value mapped to the given key at the specified index in the store
func (s *OHLCSStore) SetIdx(key string, idx int, value *OHLC) error {
	s.Lock()
	v := s.setIdx(key, idx, value)
	s.Unlock()

	return v
}

func (s *OHLCSStore) get(key string) ([]OHLC, bool) {
	// explicitly return second return value
	v, ok := s.store[key]

	return v, ok
}

// Get accesses the value for the given key
func (s *OHLCSStore) Get(key string) ([]OHLC, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *OHLCSStore) getIdx(key string, idx int) (OHLC, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return OHLC{}, ErrKeyDoesNotExist
	}

	// bounds check
	if idx < 0 || idx >= len(v) {
		return OHLC{}, ErrIdxOutOfBounds
	}

	return s.store[key][idx], nil
}

// GetIdx accesses the value for the given key at the specified index
func (s *OHLCSStore) GetIdx(key string, idx int) (OHLC, error) {
	s.Lock()
	v, err := s.getIdx(key, idx)
	s.Unlock()

	return v, err
}

func (s *OHLCSStore) getRange(key string, lower, upper int) ([]OHLC, error) {
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
func (s *OHLCSStore) GetRange(key string, lower, upper int) ([]OHLC, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *OHLCSStore) size() int {
	return len(s.store)
}

// Size returns the current size of the store
// Note: this is NOT capacity
func (s *OHLCSStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *OHLCSStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

// Members returns all member keys of the store
func (s *OHLCSStore) Members() []string {
	s.Lock()
	v := s.members()
	s.Unlock()

	return v
}

func (s *OHLCSStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

// IsMember checks if the given key exists in the store
func (s *OHLCSStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *OHLCSStore) memberLen(key string) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	return len(v), nil
}

// MemberLen returns the length of the series value stored at the given key
func (s *OHLCSStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *OHLCSStore) clear() {
	s.store = make(map[string][]OHLC)
}

// Clear deletes all keys in the store
func (s *OHLCSStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
