package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBoolSet(t *testing.T) {
	s := NewBoolStore()

	s.Set("foo", true)
	assert.True(t, s.store["foo"])
}

func TestBoolGet(t *testing.T) {
	s := NewBoolStore()

	// no key yet
	v, ok := s.Get("foo")
	assert.False(t, ok)

	// set key
	s.store["foo"] = true
	v, ok = s.Get("foo")
	assert.True(t, ok)
	assert.True(t, v)
}

func TestBoolSize(t *testing.T) {
	s := NewBoolStore()

	// no keys
	size := s.Size()
	assert.Equal(t, 0, size)

	// add two keys
	s.store["a"] = true
	s.store["b"] = true

	size = s.Size()
	assert.Equal(t, 2, size)
}

func TestBoolMembers(t *testing.T) {
	s := NewBoolStore()

	// no keys
	mems := s.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	s.store["a"] = true
	s.store["b"] = true

	mems = s.Members()
	assert.Equal(t, 2, len(mems))
}

func TestBoolIsMember(t *testing.T) {
	s := NewBoolStore()

	// no keys
	ok := s.IsMember("foo")
	assert.False(t, ok)

	// add key
	s.store["foo"] = true

	ok = s.IsMember("foo")
	assert.True(t, ok)
}

func TestBoolClear(t *testing.T) {
	s := NewBoolStore()

	s.store["foo"] = true
	assert.Equal(t, 1, len(s.store))

	s.Clear()
	assert.Equal(t, 0, len(s.store))
}

func TestBoolConcurrentGetAndSet(t *testing.T) {
	s := NewBoolStore()

	go func() {
		for i := 0; i < 100; i++ {
			s.Set("foo", true)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			s.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
