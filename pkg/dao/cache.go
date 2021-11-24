package dao

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

func NewCache(db Backend, bucket string) Cache {
	return &cache{db: db, bucket: bucket}
}

func (c *cache) ListAll() ([][]byte, error) {
	list, _ := c.db.List(c.bucket)
	return list, nil
}

//func (c *cache) ListAll(objs interface{}) (err error) {
//	switch reflect.TypeOf(objs).Kind() {
//	case reflect.Slice:
//		rv := reflect.ValueOf(objs)
//		var out []interface{}
//		list, _ := c.db.List(c.bucket)
//		for i := 0; i < rv.Len(); i++ {
//			out = append(out, rv.Index(i).Interface())
//		}
//		objs = out
//
//		for _, item := range list {
//			var f interface{}
//			_ = json.Unmarshal(item, &f)
//			rv = append(rv, &f)
//		}
//	default:
//		return errors.New("only slice support")
//	}
//
//	return nil
//}

func (c *cache) Find(identity string) (interface{}, error) {
	panic("implement me")
}

func (c *cache) Invalid() error {
	panic("implement me")
}

func (c *cache) IsExpired() bool {
	panic("implement me")
}
