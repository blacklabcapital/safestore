package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInt32Set(t *testing.T) {
	s := NewInt32Store()

	s.Set("foo", 10)
	assert.Equal(t, int32(10), s.store["foo"])
}

func TestInt32Get(t *testing.T) {
	bs := NewInt32Store()

	// no key yet
	v, ok := bs.Get("foo")
	assert.False(t, ok)

	// set key
	bs.store["foo"] = 10
	v, ok = bs.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, int32(10), v)
}

func TestInt32Size(t *testing.T) {
	bs := NewInt32Store()

	// no keys
	size := bs.Size()
	assert.Equal(t, 0, size)

	// add two keys
	bs.store["a"] = 10
	bs.store["b"] = 11

	size = bs.Size()
	assert.Equal(t, 2, size)
}

func TestInt32Members(t *testing.T) {
	bs := NewInt32Store()

	// no keys
	mems := bs.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	bs.store["a"] = 10
	bs.store["b"] = 11

	mems = bs.Members()
	assert.Equal(t, 2, len(mems))
}

func TestInt32IsMember(t *testing.T) {
	bs := NewInt32Store()

	// no keys
	ok := bs.IsMember("foo")
	assert.False(t, ok)

	// add key
	bs.store["foo"] = 10

	ok = bs.IsMember("foo")
	assert.True(t, ok)
}

func TestInt32Clear(t *testing.T) {
	bs := NewInt32Store()

	bs.store["foo"] = 10
	assert.Equal(t, 1, len(bs.store))

	bs.Clear()
	assert.Equal(t, 0, len(bs.store))
}

func TestInt32ConcurrentGetAndSet(t *testing.T) {
	bs := NewInt32Store()

	go func() {
		for i := 0; i < 100; i++ {
			bs.Set("foo", 10)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			bs.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
