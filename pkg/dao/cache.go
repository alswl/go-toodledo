package dao

// Cache ...
type Cache interface {
	ListAll() ([][]byte, error)
	Find(identity string) (interface{}, error)
	Invalid() error
	IsExpired() bool
}

type cache struct {
	db     Backend
	bucket string
}

// NewCache ...
func NewCache(db Backend, bucket string) Cache {
	return &cache{db: db, bucket: bucket}
}

// ListAll ...
func (c *cache) ListAll() ([][]byte, error) {
	list, err := c.db.List(c.bucket)
	return list, err
}

// Find ...
func (c *cache) Find(identity string) (interface{}, error) {
	panic("implement me")
}

// Invalid ...
func (c *cache) Invalid() error {
	panic("implement me")
}

// IsExpired ...
func (c *cache) IsExpired() bool {
	panic("implement me")
}
