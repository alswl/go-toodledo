package dao

import (
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

func NewBoltDB(config common.ToodledoConfig) *bolt.DB {
	db, err := bolt.Open(config.Database.DataFile, 0600, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
