package seriesstore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockIntSeries() []int {
	return []int{1, 2, 3, 4, 5}
}

func TestIntSet(t *testing.T) {
	ss := NewIntSStore()

	ss.Set("foo", mockIntSeries())
	assert.Equal(t, ss.store["foo"], mockIntSeries())
}

func TestIntSetIdx(t *testing.T) {
	// key not exist
	ss := NewIntSStore()

	err := ss.SetIdx("foo", 1, 10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockIntSeries()
	err = ss.SetIdx("foo", 1, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, ss.store["foo"][1])

	// last idx
	err = ss.SetIdx("foo", 4, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, ss.store["foo"][1])

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

func TestIntGet(t *testing.T) {
	ss := NewIntSStore()

	// no key yet
	series, ok := ss.Get("foo")
	assert.False(t, ok)

	// set key
	ss.store["foo"] = mockIntSeries()
	series, ok = ss.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, series, mockIntSeries())
}

func TestIntGetIdx(t *testing.T) {
	ss := NewIntSStore()

	// no key
	v, err := ss.GetIdx("foo", 1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockIntSeries()
	v, err = ss.GetIdx("foo", 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	// last idx
	v, err = ss.GetIdx("foo", 4)
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

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

func TestIntGetRange(t *testing.T) {
	ss := NewIntSStore()

	// no key
	rng, err := ss.GetRange("foo", 0, 5)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockIntSeries()

	// full range
	rng, err = ss.GetRange("foo", 0, 5)
	assert.Nil(t, err)
	assert.Equal(t, mockIntSeries(), rng)

	// partial range
	rng, err = ss.GetRange("foo", 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3}, rng)

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

func TestIntSize(t *testing.T) {
	ss := NewIntSStore()

	// no keys
	size := ss.Size()
	assert.Equal(t, 0, size)

	// add two keys
	ss.store["a"] = mockIntSeries()
	ss.store["b"] = mockIntSeries()

	size = ss.Size()
	assert.Equal(t, 2, size)
}

func TestIntMembers(t *testing.T) {
	ss := NewIntSStore()

	// no keys
	mems := ss.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	ss.store["a"] = mockIntSeries()
	ss.store["b"] = mockIntSeries()

	mems = ss.Members()
	assert.Equal(t, 2, len(mems))
}

func TestIntIsMember(t *testing.T) {
	ss := NewIntSStore()

	// no keys
	ok := ss.IsMember("foo")
	assert.False(t, ok)

	// add key
	ss.store["foo"] = mockIntSeries()

	ok = ss.IsMember("foo")
	assert.True(t, ok)
}

func TestIntMemberLen(t *testing.T) {
	ss := NewIntSStore()

	// no keys
	length, err := ss.MemberLen("foo")
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockIntSeries()

	length, err = ss.MemberLen("foo")
	assert.Nil(t, err)
	assert.Equal(t, 5, length)
}

func TestIntClear(t *testing.T) {
	ss := NewIntSStore()

	ss.store["foo"] = mockIntSeries()
	assert.Equal(t, 1, len(ss.store))

	ss.Clear()
	assert.Equal(t, 0, len(ss.store))
}

func TestIntConcurrentGetAndSet(t *testing.T) {
	ss := NewIntSStore()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Set("foo", mockIntSeries())
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
