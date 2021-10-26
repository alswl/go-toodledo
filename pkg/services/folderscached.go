package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

type folderCachedService struct {
	*folderService
	db         *bolt.DB
	accountSvc AccountService
}

func NewFolderCachedService(svc0 *folderService, accountSvc AccountService, db *bolt.DB) FolderService {
	s := folderCachedService{folderService: svc0, accountSvc: accountSvc, db: db}
	return &s
}

func (s *folderCachedService) listAll() ([]*models.Folder, error) {
	var fs []*models.Folder
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("folders"))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var f models.Folder
			_ = json.Unmarshal(v, &f)
			fs = append(fs, &f)
		}
		return nil
	})
	return fs, nil
}

func (s *folderCachedService) put2DB(folders []*models.Folder) error {
	s.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("folders"))
		for _, f := range folders {
			bytes, _ := json.Marshal(f)
			b.Put(([]byte)(f.Name), bytes)
		}
		return nil
	})
	return nil
}

func (s *folderCachedService) folderIsExpired() bool {
	// XXX if expired
	var meCached models.Account
	s.db.Update(func(tx *bolt.Tx) error {
		_, _ = tx.CreateBucketIfNotExists([]byte("auth"))
		return nil
	})
	s.db.View(func(tx *bolt.Tx) error {
		// XXX move to auth
		b := tx.Bucket([]byte("auth"))
		c := b.Get([]byte("me"))
		_ = json.Unmarshal(c, &meCached)
		return nil
	})

	me, err := s.accountSvc.FindMe()
	s.db.Update(func(tx *bolt.Tx) error {
		// XXX move to auth
		b := tx.Bucket([]byte("auth"))
		json, _ := json.Marshal(me)
		b.Put([]byte("me"), json)
		return nil
	})

	// XXX save to cache
	if err != nil {
		logrus.WithField("me", me).WithError(err).Error("request failed")
		return true
	}

	return me.LasteditFolder <= meCached.LasteditFolder
}

func (s *folderCachedService) syncIfExpired() error {
	if s.folderIsExpired() {
		return nil
	}

	all, err := s.folderService.ListAll()
	if err != nil {
		return err
	}
	s.put2DB(all)
	return nil
}

func (s *folderCachedService) ListAll() ([]*models.Folder, error) {
	err := s.syncIfExpired()
	if err != nil {
		return nil, err
	}

	return s.listAll()
}

func (s *folderCachedService) FindByName(name string) (*models.Folder, error) {
	err := s.syncIfExpired()
	if err != nil {
		return nil, err
	}

	return s.findByName(name)
}

func (s *folderCachedService) findByName(name string) (*models.Folder, error) {
	var f *models.Folder
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("folders"))
		if b == nil {
			return nil
		}
		c := b.Get([]byte(name))
		_ = json.Unmarshal(c, &f)
		return nil
	})
	// TODO nil
	return f, nil
}
