package dal

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/pkg/errors"
	boltdb "go.etcd.io/bbolt"
)

const perm0755 = 0755
const perm0600 = 0600

var instance Backend

var once sync.Once

// bolt is port from https://github.com/alibaba/pouch/blob/master/pkg/meta/boltdb.go
type bolt struct {
	db *boltdb.DB
	sync.Mutex
}

// NewBoltDB is used to make bolt metadata store instance.
func NewBoltDB(config models.ToodledoConfigDatabase) (Backend, error) {
	// opt := &boltdb.Options{
	//	Timeout: time.Second * 10,
	//}

	dir := filepath.Dir(config.DataFile)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, perm0755); err != nil {
			return nil, errors.Wrapf(err, "create metadata path, %s", dir)
		}
	}

	b := &bolt{}

	db, err := boltdb.Open(config.DataFile, perm0600, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "open boltdb, %s", config.DataFile)
	}
	for _, bucket := range config.Buckets {
		if err = b.prepare(db, []byte(bucket)); err != nil {
			return nil, err
		}
	}
	b.db = db

	return b, nil
}

func ProvideBackend(config models.ToodledoConfigDatabase) (Backend, error) {
	var err error
	once.Do(func() {
		instance, err = NewBoltDB(config)
		if err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (b *bolt) prepare(db *boltdb.DB, bucket []byte) error {
	b.Lock()
	defer b.Unlock()

	return db.Update(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return errors.Wrap(err, "create bucket in boltdb")
		}
		return nil
	})
}

// Path returns boltdb store file.
func (b *bolt) Path(key string) string {
	return b.db.Path()
}

// Keys return all keys for boltdb.
func (b *bolt) Keys(bucket string) ([]string, error) {
	keys := make([]string, 0)

	b.Lock()
	defer b.Unlock()

	err := b.db.View(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}

		return bkt.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	})

	return keys, err
}

// Put is used to put metadate into boltdb.
func (b *bolt) Put(bucket, key string, value []byte) error {
	b.Lock()
	defer b.Unlock()

	return b.db.Update(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}
		if err := bkt.Put([]byte(key), value); err != nil {
			return errors.Wrapf(err, "put key %s in boltdb", key)
		}
		return nil
	})
}

// Del is used to delete metadate from boltdb.
func (b *bolt) Remove(bucket string, key string) error {
	b.Lock()
	defer b.Unlock()

	return b.db.Update(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}
		return bkt.Delete([]byte(key))
	})
}

// Get returns metadata from boltdb.
func (b *bolt) Get(bucket string, key string) ([]byte, error) {
	var value []byte

	b.Lock()
	defer b.Unlock()

	err := b.db.View(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}
		if value = bkt.Get([]byte(key)); value == nil {
			return ErrObjectNotFound
		}
		return nil
	})

	return value, err
}

// List returns all metadata in boltdb.
func (b *bolt) List(bucket string) ([][]byte, error) {
	const defSize = 20
	values := make([][]byte, 0, defSize)

	b.Lock()
	defer b.Unlock()

	err := b.db.View(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}

		return bkt.ForEach(func(k, v []byte) error {
			values = append(values, v)
			return nil
		})
	})

	return values, err
}

// Truncate ...
func (b *bolt) Truncate(bucket string) error {
	b.Lock()
	err := b.db.Update(func(tx *boltdb.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt == nil {
			return ErrBucketNotFound
		}

		err := tx.DeleteBucket([]byte(bucket))
		if err != nil {
			return errors.Wrapf(err, "truncate %s in boltdb", bucket)
		}
		return nil
	})
	b.Unlock()
	if err != nil {
		return err
	}
	return b.prepare(b.db, []byte(bucket))
}

// Close releases all database resources.
// All transactions must be closed before closing the database.
func (b *bolt) Close() error {
	return b.db.Close()
}
