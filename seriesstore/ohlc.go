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

// Implements the SeriesStore interface
// A store of OHLC slices
// All get and set functions provide bound checks
// Embedded sync.Mutex to provide atomic operation ability
type OHLCSStore struct {
	sync.Mutex
	store map[string][]OHLC
}

func NewOHLCSStore() *OHLCSStore {
	return &OHLCSStore{store: make(map[string][]OHLC)}
}

func (s *OHLCSStore) set(key string, value []OHLC) {
	s.store[key] = value
}

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

func (s *OHLCSStore) GetRange(key string, lower, upper int) ([]OHLC, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *OHLCSStore) size() int {
	return len(s.store)
}

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

func (s *OHLCSStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *OHLCSStore) clear() {
	s.store = make(map[string][]OHLC)
}

func (s *OHLCSStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
