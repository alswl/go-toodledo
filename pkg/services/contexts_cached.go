package services

import (
	"encoding/json"
	"sync"

	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/thoas/go-funk"
)

var ContextBucket = "contexts"

type contextCachedService struct {
	syncLock sync.Mutex

	svc        ContextService
	db         dal.Backend
	accountSvc AccountService
}

var contextLocalServiceInstance ContextPersistenceService
var contextLocalServiceOnce sync.Once

// NewContextCachedService ...
func NewContextCachedService(contextsvc ContextService, accountSvc AccountService,
	db dal.Backend) ContextPersistenceService {
	s := contextCachedService{
		svc:        contextsvc,
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

func ProvideContextCachedService(svc ContextService, accountSvc AccountService,
	db dal.Backend) ContextPersistenceService {
	if contextLocalServiceInstance == nil {
		contextLocalServiceOnce.Do(func() {
			contextLocalServiceInstance = NewContextCachedService(svc, accountSvc, db)
		})
	}
	return contextLocalServiceInstance
}

func (s *contextCachedService) Sync() error {
	all, err := s.svc.ListAll()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()
	err = s.Clean()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		_ = s.db.Put(ContextBucket, f.Name, bytes)
	}
	return nil
}

func (s *contextCachedService) PartialSync(lastEditTime *int32) error {
	return s.Sync()
}

// Rename ...
func (s *contextCachedService) Rename(name string, newName string) (*models.Context, error) {
	_ = s.Clean()
	return s.svc.Rename(name, newName)
}

func (s *contextCachedService) Archive(id int, isArchived bool) (*models.Context, error) {
	// TODO implement me
	panic("implement me")
}

// Delete ...
func (s *contextCachedService) Delete(name string) error {
	_ = s.Clean()
	return s.svc.Delete(name)
}

// Create ...
func (s *contextCachedService) Create(name string) (*models.Context, error) {
	_ = s.Clean()
	return s.svc.Create(name)
}

// ListAll ...
func (s *contextCachedService) ListAll() ([]*models.Context, error) {
	fs := make([]*models.Context, 0)
	all, err := s.db.List(ContextBucket)
	if err != nil {
		return nil, err
	}
	for _, bytes := range all {
		f := &models.Context{}
		_ = json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

func (s *contextCachedService) FindByID(id int64) (*models.Context, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Context) bool {
		return x.ID == id
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

// Find ...
func (s *contextCachedService) Find(name string) (*models.Context, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

// LocalClear ...
func (s *contextCachedService) Clean() error {
	err := s.db.Truncate(ContextBucket)
	if err != nil {
		return err
	}
	return nil
}
