package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dao"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type FolderCachedService interface {
	Invalidate() error

	Find(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
	Rename(name string, newName string) (*models.Folder, error)
	ArchiveFolder(id int, isArchived bool) (*models.Folder, error)
	Delete(name string) error
	Create(name string) (*models.Folder, error)
}

type folderCachedService struct {
	svc        FolderService
	cache      dao.Cache
	db         dao.Backend
	accountSvc AccountService
}

func NewFolderCachedService(folderSvc FolderService, accountSvc AccountService, db dao.Backend) FolderCachedService {
	s := folderCachedService{
		svc:        folderSvc,
		cache:      dao.NewCache(db, "folders"),
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

func (s *folderCachedService) Rename(name string, newName string) (*models.Folder, error) {
	s.Invalidate()
	return s.svc.Rename(name, newName)
}

func (s *folderCachedService) ArchiveFolder(id int, isArchived bool) (*models.Folder, error) {
	s.Invalidate()
	return s.svc.ArchiveFolder(id, isArchived)
}

func (s *folderCachedService) Delete(name string) error {
	s.Invalidate()
	return s.svc.Delete(name)
}

func (s *folderCachedService) Create(name string) (*models.Folder, error) {
	s.Invalidate()
	return s.svc.Create(name)
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
	// TODO ticker
	if !s.folderIsExpired() {
		return nil
	}

	logrus.Debug("folder is not expired")
	return s.sync()
}

func (s *folderCachedService) sync() error {
	all, err := s.svc.ListAll()
	if err != nil {
		return err
	}
	err = s.Invalidate()
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
	fs := make([]*models.Folder, 0)
	all, err := s.cache.ListAll()
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		s.syncIfExpired()
		all, err = s.cache.ListAll()
		if err != nil {
			return nil, err
		}
	}
	for _, bytes := range all {
		f := &models.Folder{}
		json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

func (s *folderCachedService) Find(name string) (*models.Folder, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Folder) bool {
		return x.Name == name
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *folderCachedService) Invalidate() error {
	err := s.db.Truncate("folders")
	if err != nil {
		return err
	}
	return nil
}
