package seriesstore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockFloat32Series() []float32 {
	return []float32{1.0, 2.0, 3.0, 4.0, 5.0}
}

func TestFloat32Set(t *testing.T) {
	ss := NewFloat32SStore()

	ss.Set("foo", mockFloat32Series())
	assert.Equal(t, ss.store["foo"], mockFloat32Series())
}

func TestFloat32SetIdx(t *testing.T) {
	// key not exist
	ss := NewFloat32SStore()

	err := ss.SetIdx("foo", 1, 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockFloat32Series()
	err = ss.SetIdx("foo", 1, 10.0)
	assert.Nil(t, err)
	assert.Equal(t, float32(10.0), ss.store["foo"][1])

	// last idx
	err = ss.SetIdx("foo", 4, 10.0)
	assert.Nil(t, err)
	assert.Equal(t, float32(10.0), ss.store["foo"][4])

	// out of bounds
	// lower
	err = ss.SetIdx("foo", -1, 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)

	// upper
	err = ss.SetIdx("foo", 5, 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)
}

func TestFloat32Get(t *testing.T) {
	ss := NewFloat32SStore()

	// no key yet
	series, ok := ss.Get("foo")
	assert.False(t, ok)

	// set key
	ss.store["foo"] = mockFloat32Series()
	series, ok = ss.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, series, mockFloat32Series())
}

func TestFloat32GetIdx(t *testing.T) {
	ss := NewFloat32SStore()

	// no key
	v, err := ss.GetIdx("foo", 1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockFloat32Series()
	v, err = ss.GetIdx("foo", 0)
	assert.Nil(t, err)
	assert.Equal(t, float32(1), v)

	// last idx
	v, err = ss.GetIdx("foo", 4)
	assert.Nil(t, err)
	assert.Equal(t, float32(5.0), v)

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

func TestFloat32GetRange(t *testing.T) {
	ss := NewFloat32SStore()

	// no key
	rng, err := ss.GetRange("foo", 0, 5)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockFloat32Series()

	// full range
	rng, err = ss.GetRange("foo", 0, 5)
	assert.Nil(t, err)
	assert.Equal(t, mockFloat32Series(), rng)

	// partial range
	rng, err = ss.GetRange("foo", 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, []float32{1, 2, 3}, rng)

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

func TestFloat32Size(t *testing.T) {
	ss := NewFloat32SStore()

	// no keys
	size := ss.Size()
	assert.Equal(t, 0, size)

	// add two keys
	ss.store["a"] = mockFloat32Series()
	ss.store["b"] = mockFloat32Series()

	size = ss.Size()
	assert.Equal(t, 2, size)
}

func TestFloat32Members(t *testing.T) {
	ss := NewFloat32SStore()

	// no keys
	mems := ss.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	ss.store["a"] = mockFloat32Series()
	ss.store["b"] = mockFloat32Series()

	mems = ss.Members()
	assert.Equal(t, 2, len(mems))
}

func TestFloat32IsMember(t *testing.T) {
	ss := NewFloat32SStore()

	// no keys
	ok := ss.IsMember("foo")
	assert.False(t, ok)

	// add key
	ss.store["foo"] = mockFloat32Series()

	ok = ss.IsMember("foo")
	assert.True(t, ok)
}

func TestFloat32MemberLen(t *testing.T) {
	ss := NewFloat32SStore()

	// no keys
	length, err := ss.MemberLen("foo")
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockFloat32Series()

	length, err = ss.MemberLen("foo")
	assert.Nil(t, err)
	assert.Equal(t, 5, length)
}

func TestFloat32Clear(t *testing.T) {
	ss := NewFloat32SStore()

	ss.store["foo"] = mockFloat32Series()
	assert.Equal(t, 1, len(ss.store))

	ss.Clear()
	assert.Equal(t, 0, len(ss.store))
}

func TestFloat32ConcurrentGetAndSet(t *testing.T) {
	ss := NewFloat32SStore()

	// set
	go func() {
		for i := 0; i < 100; i++ {
			ss.Set("foo", mockFloat32Series())
		}
	}()

	// get
	go func() {
		for i := 0; i < 100; i++ {
			ss.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
