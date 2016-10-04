package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInt64Set(t *testing.T) {
	s := NewInt64Store()

	s.Set("foo", 10)
	assert.Equal(t, int64(10), s.store["foo"])
}

func TestInt64Get(t *testing.T) {
	s := NewInt64Store()

	// no key yet
	v, ok := s.Get("foo")
	assert.False(t, ok)

	// set key
	s.store["foo"] = 10
	v, ok = s.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, int64(10), v)
}

func TestInt64Size(t *testing.T) {
	s := NewInt64Store()

	// no keys
	size := s.Size()
	assert.Equal(t, 0, size)

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	size = s.Size()
	assert.Equal(t, 2, size)
}

func TestInt64Members(t *testing.T) {
	s := NewInt64Store()

	// no keys
	mems := s.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	mems = s.Members()
	assert.Equal(t, 2, len(mems))
}

func TestInt64IsMember(t *testing.T) {
	s := NewInt64Store()

	// no keys
	ok := s.IsMember("foo")
	assert.False(t, ok)

	// add key
	s.store["foo"] = 10

	ok = s.IsMember("foo")
	assert.True(t, ok)
}

func TestInt64Clear(t *testing.T) {
	s := NewInt64Store()

	s.store["foo"] = 10
	assert.Equal(t, 1, len(s.store))

	s.Clear()
	assert.Equal(t, 0, len(s.store))
}

func TestInt64ConcurrentGetAndSet(t *testing.T) {
	s := NewInt64Store()

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
