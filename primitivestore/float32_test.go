package primitivestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFloat32Set(t *testing.T) {
	s := NewFloat32Store()

	s.Set("foo", 10.5)
	assert.Equal(t, float32(10.5), s.store["foo"])
}

func TestFloat32Get(t *testing.T) {
	s := NewFloat32Store()

	// no key yet
	v, ok := s.Get("foo")
	assert.False(t, ok)

	// set key
	s.store["foo"] = 10.5
	v, ok = s.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, float32(10.5), v)
}

func TestFloat32Size(t *testing.T) {
	s := NewFloat32Store()

	// no keys
	size := s.Size()
	assert.Equal(t, 0, size)

	// add two keys
	s.store["a"] = 10.5
	s.store["b"] = 11.5

	size = s.Size()
	assert.Equal(t, 2, size)
}

func TestFloat32Members(t *testing.T) {
	s := NewFloat32Store()

	// no keys
	mems := s.Members()
	assert.Equal(t, 0, len(mems))

	// add two keys
	s.store["a"] = 10.5
	s.store["b"] = 11.5

	mems = s.Members()
	assert.Equal(t, 2, len(mems))
}

func TestFloat32IsMember(t *testing.T) {
	s := NewFloat32Store()

	// no keys
	ok := s.IsMember("foo")
	assert.False(t, ok)

	// add key
	s.store["foo"] = 10.5

	ok = s.IsMember("foo")
	assert.True(t, ok)
}

func TestFloat32Clear(t *testing.T) {
	s := NewFloat32Store()

	s.store["foo"] = 10.5
	assert.Equal(t, 1, len(s.store))

	s.Clear()
	assert.Equal(t, 0, len(s.store))
}

func TestFloat32ConcurrentGetAndSet(t *testing.T) {
	s := NewFloat32Store()

	go func() {
		for i := 0; i < 100; i++ {
			s.Set("foo", 10.5)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			s.Get("foo")
		}
	}()

	time.Sleep(time.Second * 2)
}
