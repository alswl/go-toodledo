package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"sync"
)

// FolderCachedService ...
// TODO remove this, using syncer
type FolderCachedService interface {
	LocalTruncate() error
	Sync() error

	Find(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
	Rename(name string, newName string) (*models.Folder, error)
	Archive(id int, isArchived bool) (*models.Folder, error)
	Delete(name string) error
	Create(name string) (*models.Folder, error)
}

type folderCachedService struct {
	syncLock sync.Mutex

	svc        FolderService
	cache      dal.Cache
	db         dal.Backend
	accountSvc AccountService
}

// NewFolderCachedService ...
func NewFolderCachedService(folderSvc FolderService, accountSvc AccountService, db dal.Backend) FolderCachedService {
	s := folderCachedService{
		svc:        folderSvc,
		cache:      dal.NewCache(db, "folders"),
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

// Rename ...
func (s *folderCachedService) Rename(name string, newName string) (*models.Folder, error) {
	s.LocalTruncate()
	return s.svc.Rename(name, newName)
}

// Archive ...
func (s *folderCachedService) Archive(id int, isArchived bool) (*models.Folder, error) {
	s.LocalTruncate()
	return s.svc.Archive(id, isArchived)
}

// Delete ...
func (s *folderCachedService) Delete(name string) error {
	s.LocalTruncate()
	return s.svc.Delete(name)
}

// Create ...
func (s *folderCachedService) Create(name string) (*models.Folder, error) {
	s.LocalTruncate()
	return s.svc.Create(name)
}

func (s *folderCachedService) folderIsExpired() bool {
	var meCached models.Account
	// FIXME userService
	c, err := s.db.Get("auth", "me")
	if err == dal.ErrObjectNotFound {
		// missing
		me, err := s.accountSvc.Me()
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

func (s *folderCachedService) Sync() error {
	all, err := s.svc.ListAll()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()
	err = s.LocalTruncate()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		s.db.Put("folders", f.Name, bytes)
	}
	return nil
}

// ListAll ...
func (s *folderCachedService) ListAll() ([]*models.Folder, error) {
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

// Find ...
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

func (s *folderCachedService) LocalTruncate() error {
	err := s.db.Truncate("folders")
	if err != nil {
		return err
	}
	return nil
}
