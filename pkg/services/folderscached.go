package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/dao"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
)

type folderCachedService struct {
	*folderService
	cache      dao.Cache
	db         dao.Backend
	accountSvc AccountService
}

func NewFolderCachedService(folderSvc *folderService, accountSvc AccountService, db dao.Backend) FolderService {
	s := folderCachedService{
		folderService: folderSvc,
		cache:         dao.NewCache(db, "folders"),
		db:            db,
		accountSvc:    accountSvc,
	}
	return &s
}

func (s *folderCachedService) listAll() ([]*models.Folder, error) {
	var objs []interface{}
	err := s.cache.ListAll(objs)
	if err != nil {
		return nil, err
	}
	fs := make([]*models.Folder, len(objs))
	for _, obj := range objs {
		fs = append(fs, obj.(*models.Folder))
	}
	return fs, nil
}

func (s *folderCachedService) folderIsExpired() bool {
	// XXX if expired
	var meCached models.Account
	//s.db.Update(func(tx *bolt.Tx) error {
	//	_, _ = tx.CreateBucketIfNotExists([]byte("auth"))
	//	return nil
	//})
	c, _ := s.db.Get("auth", "me")
	_ = json.Unmarshal(c, &meCached)
	//s.db.View(func(tx *bolt.Tx) error {
	//	// XXX move to auth
	//	b := tx.Bucket([]byte("auth"))
	//	c := b.Get([]byte("me"))
	//	_ = json.Unmarshal(c, &meCached)
	//	return nil
	//})

	me, err := s.accountSvc.FindMe()
	json, _ := json.Marshal(me)
	s.db.Put("auth", "me", json)
	//s.db.Update(func(tx *bolt.Tx) error {
	//	// XXX move to auth
	//	b := tx.Bucket([]byte("auth"))
	//	json, _ := json.Marshal(me)
	//	b.Put([]byte("me"), json)
	//	return nil
	//})

	// XXX save to cache
	if err != nil {
		logrus.WithField("me", me).WithError(err).Error("request failed")
		return true
	}

	return me.LasteditFolder <= meCached.LasteditFolder
}

func (s *folderCachedService) syncIfExpired() error {
	if !s.folderIsExpired() {
		return nil
	}

	logrus.Debug("folder is not expired")
	return s.sync()
}

func (s *folderCachedService) sync() error {
	all, err := s.folderService.ListAll()
	if err != nil {
		return err
	}
	err = s.db.Truncate("folders")
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		s.db.Put("folders", f.Name, bytes)
	}
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
	c, _ := s.db.Get("folders", name)
	_ = json.Unmarshal(c, &f)
	//s.db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("folders"))
	//	if b == nil {
	//		return nil
	//	}
	//	c := b.Get([]byte(name))
	//	_ = json.Unmarshal(c, &f)
	//	return nil
	//})
	// TODO nil
	return f, nil
}
