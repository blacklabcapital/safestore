package seriesstore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockUint64Series() []uint64 {
	return []uint64{1, 2, 3, 4, 5}
}

func TestUint64Set(t *testing.T) {
	ss := NewUint64SStore()

	ss.Set("foo", mockUint64Series())
	assert.Equal(t, ss.store["foo"], mockUint64Series())
}

func TestUint64SetIdx(t *testing.T) {
	// key not exist
	ss := NewUint64SStore()

	err := ss.SetIdx("foo", 1, 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockUint64Series()
	err = ss.SetIdx("foo", 1, 10)
	assert.Nil(t, err)
	assert.Equal(t, uint64(10), ss.store["foo"][1])

	// last idx
	err = ss.SetIdx("foo", 4, 10)
	assert.Nil(t, err)
	assert.Equal(t, uint64(10), ss.store["foo"][1])

	// out of bounds
	// lower
	err = ss.SetIdx("foo", -1, 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)

	// upper
	err = ss.SetIdx("foo", 5, 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)
}

func TestUint64Get(t *testing.T) {
	ss := NewUint64SStore()

	// no key yet
	series, ok := ss.Get("foo")
	assert.False(t, ok)

	// set key
	ss.store["foo"] = mockUint64Series()
	series, ok = ss.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, series, mockUint64Series())
}

func TestUint64GetIdx(t *testing.T) {
	ss := NewUint64SStore()

	// no key
	v, err := ss.GetIdx("foo", 1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockUint64Series()
	v, err = ss.GetIdx("foo", 0)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), v)

	// last idx
	v, err = ss.GetIdx("foo", 4)
	assert.Nil(t, err)
	assert.Equal(t, uint64(5), v)

	// out of bounds
	// lower
	v, err = ss.GetIdx("foo", -1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)

	// upper
	v, err = ss.GetIdx("foo", 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)
}

func TestUint64GetRange(t *testing.T) {
	ss := NewUint64SStore()

	// no key
	rng, err := ss.GetRange("foo", 0, 5)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockUint64Series()

	// full range
	rng, err = ss.GetRange("foo", 0, 5)
	assert.Nil(t, err)
	assert.Equal(t, mockUint64Series(), rng)

	// partial range
	rng, err = ss.GetRange("foo", 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, []uint64{1, 2, 3}, rng)

	// out of bounds
	// lower
	rng, err = ss.GetRange("foo", -1, 3)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)

	// upper
	rng, err = ss.GetRange("foo", 0, 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)
}

func TestUint64Size(t *testing.T) {
	ss := NewUint64SStore()

	// no keys
	size := ss.Size()
	assert.Equal(t, 0, size)

	// add two keys
	ss.store["a"] = mockUint64Series()
	ss.store["b"] = mockUint64Series()

	size = ss.Size()
	assert.Equal(t, 2, size)
}

func TestUint64Members(t *testing.T) {
	ss := NewUint64SStore()

	// no keys
	mems := ss.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	ss.store["a"] = mockUint64Series()
	ss.store["b"] = mockUint64Series()

	mems = ss.Members()
	assert.Equal(t, 2, len(mems))
}

func TestUint64IsMember(t *testing.T) {
	ss := NewUint64SStore()

	// no keys
	ok := ss.IsMember("foo")
	assert.False(t, ok)

	// add key
	ss.store["foo"] = mockUint64Series()

	ok = ss.IsMember("foo")
	assert.True(t, ok)
}

func TestUint64MemberLen(t *testing.T) {
	ss := NewUint64SStore()

	// no keys
	length, err := ss.MemberLen("foo")
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockUint64Series()

	length, err = ss.MemberLen("foo")
	assert.Nil(t, err)
	assert.Equal(t, 5, length)
}

func TestUint64Clear(t *testing.T) {
	ss := NewUint64SStore()

	ss.store["foo"] = mockUint64Series()
	assert.Equal(t, 1, len(ss.store))

	ss.Clear()
	assert.Equal(t, 0, len(ss.store))
}

func TestUint64ConcurrentGetAndSet(t *testing.T) {
	ss := NewUint64SStore()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Set("foo", mockUint64Series())
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
