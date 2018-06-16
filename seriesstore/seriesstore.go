package seriesstore

import (
	"errors"
)

var (
	// ErrKeyDoesNotExist is thrown when a search key is not found
	ErrKeyDoesNotExist = errors.New("key does not exist")
	// ErrIdxOutOfBounds is thrown when given indices for a range are out of bounds
	ErrIdxOutOfBounds = errors.New("index out of bounds")
)

// A SeriesStore is a key/value storage that stores a data series
// Provides atomic methods safe for concurrent use for setting and getting data
// Provides index and range access methods for safely getting specific values
// from underlying array
type SeriesStore interface {
	// Set sets the key in the store to the given series value
	Set(key string, value []interface{})

	// SetIdx sets the index value of series of the given key in the store
	SetIdx(key string, idx int, value interface{}) error

	// Get gets the series value from the store for the given key
	// returns the series and boolean if key exists
	Get(key string) ([]interface{}, bool)

	// GetIdx gets the value of the series at the given index for the given key
	GetIdx(key string, idx int) (interface{}, error)

	// GetRange gets a range of values in the series of the given key from the store
	// for the given index bounds. Bounds are [Inclusive:Exclusive]
	GetRange(key string, lower, upper int) ([]interface{}, error)

	// Size returns the size of the store
	Size() int

	// Members returns a list of string keys in the store
	Members() []string

	// IsMember checks if the given key is a member of the store
	isMember(key string) bool

	// MemberLen gets the length of the series value for the given key
	MemberLen(key string) (int, error)

	// Clear deletes all stores keys and values
	Clear()
}
