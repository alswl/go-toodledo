package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/thoas/go-funk"
	"sync"
)

// GoalLocalService is a cached service
// it synced interval by fetcher
type GoalLocalService interface {
	LocalStorage
	GoalService
}

var GoalBucket = "goals"

type goalCachedService struct {
	syncLock sync.Mutex

	svc        GoalService
	db         dal.Backend
	accountSvc AccountService
}

// NewGoalLocalService ...
func NewGoalLocalService(goalsvc GoalService, accountSvc AccountService, db dal.Backend) GoalLocalService {
	s := goalCachedService{
		svc:        goalsvc,
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

func (s *goalCachedService) Sync() error {
	all, err := s.svc.ListAll()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()
	err = s.LocalClear()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		_ = s.db.Put(GoalBucket, f.Name, bytes)
	}
	return nil
}

func (s *goalCachedService) PartialSync(lastEditTime *int32) error {
	return s.Sync()
}

// Rename ...
func (s *goalCachedService) Rename(name string, newName string) (*models.Goal, error) {
	_ = s.LocalClear()
	return s.svc.Rename(name, newName)
}

// Archive ...
func (s *goalCachedService) Archive(id int, isArchived bool) (*models.Goal, error) {
	_ = s.LocalClear()
	return s.svc.Archive(id, isArchived)
}

// Delete ...
func (s *goalCachedService) Delete(name string) error {
	_ = s.LocalClear()
	return s.svc.Delete(name)
}

// Create ...
func (s *goalCachedService) Create(name string) (*models.Goal, error) {
	_ = s.LocalClear()
	return s.svc.Create(name)
}

// ListAll ...
func (s *goalCachedService) ListAll() ([]*models.Goal, error) {
	fs := make([]*models.Goal, 0)
	all, err := s.db.List(GoalBucket)
	if err != nil {
		return nil, err
	}
	for _, bytes := range all {
		f := &models.Goal{}
		_ = json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

// Find ...
func (s *goalCachedService) Find(name string) (*models.Goal, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Goal) bool {
		return x.Name == name
	}).([]*models.Goal)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *goalCachedService) FindByID(id int64) (*models.Goal, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Goal) bool {
		return x.ID == id
	}).([]*models.Goal)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *goalCachedService) LocalClear() error {
	err := s.db.Truncate(GoalBucket)
	if err != nil {
		return err
	}
	return nil
}