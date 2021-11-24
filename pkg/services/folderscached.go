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
	fs := make([]*models.Folder, 0)
	all, err := s.cache.ListAll()
	if err != nil {
		return nil, err
	}
	for _, bytes := range all {
		f := &models.Folder{}
		json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

func (s *folderCachedService) folderIsExpired() bool {
	var meCached models.Account
	// FIXME userService
	c, err := s.db.Get("auth", "me")
	if err == dao.ErrObjectNotFound {
		// missing
		me, err := s.accountSvc.FindMe()
		c, _ = json.Marshal(me)
		s.db.Put("auth", "me", c)
		// FIXME save to cache
		if err != nil {
			logrus.WithField("me", me).WithError(err).Error("request failed")
			return true
		}

	}
	_ = json.Unmarshal(c, &meCached)

	return meCached.LasteditFolder <= meCached.LasteditFolder
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

func (s *folderCachedService) Find(name string) (*models.Folder, error) {
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
