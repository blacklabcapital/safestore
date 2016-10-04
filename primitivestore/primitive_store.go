package primitivestore

// A PrimitiveStore is a key/value storage that stores primitive data type values
// Provides atomic methods safe for concurrent use for setting and getting data
type PrimitiveStore interface {
	// Set sets the key in the store to the given value
	Set(key string, value interface{})

	// Get gets the value from the store for the given key
	// returns the value and boolean if key exists
	Get(key string) ([]interface{}, bool)

	// Size returns the size of the store
	Size() int

	// Members returns a list of string keys in the store
	Members() []string

	// IsMember checks if the given key is a member of the store
	IsMember(key string) bool

	// Clear deletes all stores keys and values
	Clear()
}
