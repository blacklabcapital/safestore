package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUint32Set(t *testing.T) {
	s := NewUint32Store()

	s.Set("foo", 10)
	assert.Equal(t, uint32(10), s.store["foo"])
}

func TestUint32Get(t *testing.T) {
	s := NewUint32Store()

	// no key yet
	v, ok := s.Get("foo")
	assert.False(t, ok)

	// set key
	s.store["foo"] = 10
	v, ok = s.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, uint32(10), v)
}

func TestUint32Size(t *testing.T) {
	s := NewUint32Store()

	// no keys
	size := s.Size()
	assert.Equal(t, 0, size)

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	size = s.Size()
	assert.Equal(t, 2, size)
}

func TestUint32Members(t *testing.T) {
	s := NewUint32Store()

	// no keys
	mems := s.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	mems = s.Members()
	assert.Equal(t, 2, len(mems))
}

func TestUint32IsMember(t *testing.T) {
	s := NewUint32Store()

	// no keys
	ok := s.IsMember("foo")
	assert.False(t, ok)

	// add key
	s.store["foo"] = 10

	ok = s.IsMember("foo")
	assert.True(t, ok)
}

func TestUint32Clear(t *testing.T) {
	s := NewUint32Store()

	s.store["foo"] = 10
	assert.Equal(t, 1, len(s.store))

	s.Clear()
	assert.Equal(t, 0, len(s.store))
}

func TestUint32ConcurrentGetAndSet(t *testing.T) {
	s := NewUint32Store()

	go func() {
		for i := 0; i < 100; i++ {
			s.Set("foo", 10)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			s.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
