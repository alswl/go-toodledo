package dal

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/pkg/errors"
	boltdb "go.etcd.io/bbolt"
	"os"
	"path/filepath"
)

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

// NewBoltDB is used to make bolt metadata store instance.
func NewBoltDB(config models.ToodledoCliConfig) (Backend, error) {
	//opt := &boltdb.Options{
	//	Timeout: time.Second * 10,
	//}

	dir := filepath.Dir(config.Database.DataFile)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, errors.Wrapf(err, "create metadata path, %s", dir)
		}
	}

	b := &bolt{}

	db, err := boltdb.Open(config.Database.DataFile, 0600, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "open boltdb, %s", config.Database.DataFile)
	}
	for _, bucket := range config.Database.Buckets {
		if err := b.prepare(db, []byte(bucket)); err != nil {
			return nil, err
		}
	}
	b.db = db

	return b, nil
}
