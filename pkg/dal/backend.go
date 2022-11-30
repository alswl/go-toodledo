package dal

// Backend is an interface which describes what a store should support.
// All methods are thread-safe
// port from https://github.com/alibaba/pouch/blob/master/pkg/meta/backend.go
type Backend interface {
	// Put write key-value into store.
	Put(bucket string, key string, value []byte) error

	// Get read object from store.
	Get(bucket string, key string) ([]byte, error)

	// Remove all data of the key.
	Remove(bucket string, key string) error

	// Truncate all data of the bucket.
	Truncate(bucket string) error

	// List return all objects with specify bucket.
	List(bucket string) ([][]byte, error)

	// Keys return all keys.
	Keys(bucket string) ([]string, error)

	// Path returns the path with the specified key.
	Path(key string) string

	// Close releases all resources used by the store
	// It does not make any changes to store.
	Close() error
}
