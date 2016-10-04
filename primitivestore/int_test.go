package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	s := NewIntStore()

	s.Set("foo", 10)
	assert.Equal(t, 10, s.store["foo"])
}

func TestIntGet(t *testing.T) {
	s := NewIntStore()

	// no key yet
	v, ok := s.Get("foo")
	assert.False(t, ok)

	// set key
	s.store["foo"] = 10
	v, ok = s.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, 10, v)
}

func TestIntSize(t *testing.T) {
	s := NewIntStore()

	// no keys
	size := s.Size()
	assert.Equal(t, 0, size)

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	size = s.Size()
	assert.Equal(t, 2, size)
}

func TestIntMembers(t *testing.T) {
	s := NewIntStore()

	// no keys
	mems := s.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	s.store["a"] = 10
	s.store["b"] = 11

	mems = s.Members()
	assert.Equal(t, 2, len(mems))
}

func TestIntIsMember(t *testing.T) {
	s := NewIntStore()

	// no keys
	ok := s.IsMember("foo")
	assert.False(t, ok)

	// add key
	s.store["foo"] = 10

	ok = s.IsMember("foo")
	assert.True(t, ok)
}

func TestIntClear(t *testing.T) {
	s := NewIntStore()

	s.store["foo"] = 10
	assert.Equal(t, 1, len(s.store))

	s.Clear()
	assert.Equal(t, 0, len(s.store))
}

func TestIntConcurrentGetAndSet(t *testing.T) {
	s := NewIntStore()

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
