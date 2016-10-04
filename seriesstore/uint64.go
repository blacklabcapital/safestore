package seriesstore

import (
	"sync"
)

// Implements the SeriesStore interface
// A store of uint64 slices
// All get and set functions provide bound checks
// Embedded sync.Mutex to provide atomic operation ability
type Uint64SStore struct {
	sync.Mutex
	store map[string][]uint64
}

func NewUint64SStore() *Uint64SStore {
	return &Uint64SStore{store: make(map[string][]uint64)}
}

func (s *Uint64SStore) set(key string, value []uint64) {
	s.store[key] = value
}

func (s *Uint64SStore) Set(key string, value []uint64) {
	s.Lock()
	s.set(key, value)
	s.Unlock()
}

func (s *Uint64SStore) setIdx(key string, idx int, value uint64) error {
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

func (s *Uint64SStore) SetIdx(key string, idx int, value uint64) error {
	s.Lock()
	err := s.setIdx(key, idx, value)
	s.Unlock()

	return err
}

func (s *Uint64SStore) get(key string) ([]uint64, bool) {
	// explicitly return second return value
	v, ok := s.store[key]

	return v, ok
}

func (s *Uint64SStore) Get(key string) ([]uint64, bool) {
	s.Lock()
	v, ok := s.get(key)
	s.Unlock()

	return v, ok
}

func (s *Uint64SStore) getIdx(key string, idx int) (uint64, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	// bounds check
	if idx < 0 || idx >= len(v) {
		return 0, ErrIdxOutOfBounds
	}

	return s.store[key][idx], nil
}

func (s *Uint64SStore) GetIdx(key string, idx int) (uint64, error) {
	s.Lock()
	v, err := s.getIdx(key, idx)
	s.Unlock()

	return v, err
}

func (s *Uint64SStore) getRange(key string, lower, upper int) ([]uint64, error) {
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

func (s *Uint64SStore) GetRange(key string, lower, upper int) ([]uint64, error) {
	s.Lock()
	v, err := s.getRange(key, lower, upper)
	s.Unlock()

	return v, err
}

func (s *Uint64SStore) size() int {
	return len(s.store)
}

func (s *Uint64SStore) Size() int {
	s.Lock()
	size := s.size()
	s.Unlock()

	return size
}

func (s *Uint64SStore) members() []string {
	mems := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		mems[i] = k
		i++
	}

	return mems
}

func (s *Uint64SStore) Members() []string {
	s.Lock()
	v := s.members()
	s.Unlock()

	return v
}

func (s *Uint64SStore) isMember(key string) bool {
	_, ok := s.store[key]

	return ok
}

func (s *Uint64SStore) IsMember(key string) bool {
	s.Lock()
	ok := s.isMember(key)
	s.Unlock()

	return ok
}

func (s *Uint64SStore) memberLen(key string) (int, error) {
	v, ok := s.store[key]

	// check exists
	if !ok {
		return 0, ErrKeyDoesNotExist
	}

	return len(v), nil
}

func (s *Uint64SStore) MemberLen(key string) (int, error) {
	s.Lock()
	l, err := s.memberLen(key)
	s.Unlock()

	return l, err
}

func (s *Uint64SStore) clear() {
	s.store = make(map[string][]uint64)
}

func (s *Uint64SStore) Clear() {
	s.Lock()
	s.clear()
	s.Unlock()
}
