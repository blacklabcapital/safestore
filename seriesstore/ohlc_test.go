package seriesstore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockOHLCSeries() []OHLC {
	return []OHLC{{100.0, 200.0, 50.0, 101.0}, {101.0, 201.0, 49.0, 102.0}, {102.0, 202.0, 48.0, 103.0}}
}

func TestOHLCSet(t *testing.T) {
	ss := NewOHLCSStore()

	ss.Set("foo", mockOHLCSeries())
	assert.Equal(t, ss.store["foo"], mockOHLCSeries())
}

func TestOHLCSetIdx(t *testing.T) {
	// key not exist
	ss := NewOHLCSStore()
	candle := OHLC{109.0, 155.0, 46.0, 103.0}
	candle2 := OHLC{103.0, 159.0, 44.0, 108.0}
	err := ss.SetIdx("foo", 1, &candle)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockOHLCSeries()
	err = ss.SetIdx("foo", 1, &candle)
	assert.Nil(t, err)
	assert.Equal(t, candle, ss.store["foo"][1])

	// last idx
	err = ss.SetIdx("foo", 2, &candle2)
	assert.Nil(t, err)
	assert.Equal(t, candle2, ss.store["foo"][2])

	// out of bounds
	// lower
	err = ss.SetIdx("foo", -1, &candle)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)

	// upper
	err = ss.SetIdx("foo", 5, &candle)
	assert.NotNil(t, err)
	assert.Equal(t, ErrIdxOutOfBounds, err)
}

func TestOHLCGet(t *testing.T) {
	ss := NewOHLCSStore()

	// no key yet
	series, ok := ss.Get("foo")
	assert.False(t, ok)

	// set key
	ss.store["foo"] = mockOHLCSeries()
	series, ok = ss.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, series, mockOHLCSeries())
}

func TestOHLCGetIdx(t *testing.T) {
	ss := NewOHLCSStore()

	// no key
	v, err := ss.GetIdx("foo", 1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockOHLCSeries()
	v, err = ss.GetIdx("foo", 0)
	assert.Nil(t, err)
	assert.Equal(t, mockOHLCSeries()[0], v)

	// last idx
	v, err = ss.GetIdx("foo", 2)
	assert.Nil(t, err)
	assert.Equal(t, mockOHLCSeries()[2], v)

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

func TestOHLCGetRange(t *testing.T) {
	ss := NewOHLCSStore()

	// no key
	rng, err := ss.GetRange("foo", 0, 2)
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockOHLCSeries()

	// full range
	rng, err = ss.GetRange("foo", 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, mockOHLCSeries()[0:2], rng)

	// partial range
	rng, err = ss.GetRange("foo", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, mockOHLCSeries()[0:1], rng)

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

func TestOHLCSize(t *testing.T) {
	ss := NewOHLCSStore()

	// no keys
	size := ss.Size()
	assert.Equal(t, 0, size)

	// add two keys
	ss.store["a"] = mockOHLCSeries()
	ss.store["b"] = mockOHLCSeries()

	size = ss.Size()
	assert.Equal(t, 2, size)
}

func TestOHLCMembers(t *testing.T) {
	ss := NewOHLCSStore()

	// no keys
	mems := ss.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	ss.store["a"] = mockOHLCSeries()
	ss.store["b"] = mockOHLCSeries()

	mems = ss.Members()
	assert.Equal(t, 2, len(mems))
}

func TestOHLCIsMember(t *testing.T) {
	ss := NewOHLCSStore()

	// no keys
	ok := ss.IsMember("foo")
	assert.False(t, ok)

	// add key
	ss.store["foo"] = mockOHLCSeries()

	ok = ss.IsMember("foo")
	assert.True(t, ok)
}

func TestOHLCMemberLen(t *testing.T) {
	ss := NewOHLCSStore()

	// no keys
	length, err := ss.MemberLen("foo")
	assert.NotNil(t, err)
	assert.Equal(t, ErrKeyDoesNotExist, err)

	// add key
	ss.store["foo"] = mockOHLCSeries()

	length, err = ss.MemberLen("foo")
	assert.Nil(t, err)
	assert.Equal(t, 3, length)
}

func TestOHLCClear(t *testing.T) {
	ss := NewOHLCSStore()

	ss.store["foo"] = mockOHLCSeries()
	assert.Equal(t, 1, len(ss.store))

	ss.Clear()
	assert.Equal(t, 0, len(ss.store))
}

func TestOHLCConcurrentGetAndSet(t *testing.T) {
	ss := NewOHLCSStore()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Set("foo", mockOHLCSeries())
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ss.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
