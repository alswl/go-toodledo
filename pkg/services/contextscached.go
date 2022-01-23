package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// ContextCachedService ...
type ContextCachedService interface {
	LocalClear() error

	Find(name string) (*models.Context, error)
	ListAll() ([]*models.Context, error)
	Rename(name string, newName string) (*models.Context, error)
	Delete(name string) error
	Create(name string) (*models.Context, error)
}

type contextCachedService struct {
	svc        ContextService
	cache      dal.Cache
	db         dal.Backend
	accountSvc AccountService
}

// NewContextCachedService ...
func NewContextCachedService(contextsvc ContextService, accountSvc AccountService, db dal.Backend) ContextCachedService {
	s := contextCachedService{
		svc:        contextsvc,
		cache:      dal.NewCache(db, "contexts"),
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

// Rename ...
func (s *contextCachedService) Rename(name string, newName string) (*models.Context, error) {
	s.LocalClear()
	return s.svc.Rename(name, newName)
}

// Delete ...
func (s *contextCachedService) Delete(name string) error {
	s.LocalClear()
	return s.svc.Delete(name)
}

// Create ...
func (s *contextCachedService) Create(name string) (*models.Context, error) {
	s.LocalClear()
	return s.svc.Create(name)
}

func (s *contextCachedService) contextIsExpired() bool {
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

	return meCached.LasteditContext <= meCached.LasteditContext
}

func (s *contextCachedService) syncIfExpired() error {
	// TODO ticker
	if !s.contextIsExpired() {
		return nil
	}

	logrus.Debug("context is not expired")
	return s.sync()
}

func (s *contextCachedService) sync() error {
	all, err := s.svc.ListAll()
	if err != nil {
		return err
	}
	err = s.LocalClear()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		s.db.Put("contexts", f.Name, bytes)
	}
	return nil
}

// ListAll ...
func (s *contextCachedService) ListAll() ([]*models.Context, error) {
	fs := make([]*models.Context, 0)
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
		f := &models.Context{}
		json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

// Find ...
func (s *contextCachedService) Find(name string) (*models.Context, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

// Invalidate ...
func (s *contextCachedService) LocalClear() error {
	err := s.db.Truncate("contexts")
	if err != nil {
		return err
	}
	return nil
}
