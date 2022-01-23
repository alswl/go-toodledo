package dal

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"sync"
)

var bolts = map[string]Backend{}
var lock = sync.Mutex{}

func ProvideBackend(config models.ToodledoCliConfig) (Backend, error) {
	lock.Lock()
	defer lock.Unlock()
	if b, ok := bolts[config.Database.DataFile]; ok {
		return b, nil
	}
	b, err := NewBoltDB(config)
	if err != nil {
		return nil, err
	}
	bolts[config.Database.DataFile] = b
	return b, nil
}
