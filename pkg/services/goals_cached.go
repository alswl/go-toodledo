package services

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/thoas/go-funk"
)

const GoalsBucket = "goals"

type goalCachedService struct {
	syncLock sync.Mutex

	svc        GoalService
	db         dal.Backend
	accountSvc AccountExtService
}

func NewGoalCachedService(goalSvc GoalService, accountSvc AccountExtService, db dal.Backend) GoalPersistenceService {
	s := goalCachedService{
		svc:        goalSvc,
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
	err = s.Clean()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		_ = s.db.Put(GoalsBucket, strconv.Itoa(int(f.ID)), bytes)
	}
	return nil
}

func (s *goalCachedService) PartialSync(lastEditTime *int32) error {
	return s.Sync()
}

func (s *goalCachedService) Rename(name string, newName string) (*models.Goal, error) {
	_ = s.Clean()
	return s.svc.Rename(name, newName)
}

func (s *goalCachedService) Archive(id int, isArchived bool) (*models.Goal, error) {
	_ = s.Clean()
	return s.svc.Archive(id, isArchived)
}

func (s *goalCachedService) Delete(name string) error {
	_ = s.Clean()
	return s.svc.Delete(name)
}

func (s *goalCachedService) Create(name string) (*models.Goal, error) {
	_ = s.Clean()
	return s.svc.Create(name)
}

func (s *goalCachedService) ListAll() ([]*models.Goal, error) {
	all, err := s.ListAllWithArchived()
	if err != nil {
		return nil, err
	}
	return funk.Filter(all, func(x *models.Goal) bool {
		return x.Archived == 0
	}).([]*models.Goal), nil
}

func (s *goalCachedService) ListAllWithArchived() ([]*models.Goal, error) {
	fs := make([]*models.Goal, 0)
	all, err := s.db.List(GoalsBucket)
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

func (s *goalCachedService) Find(name string) (*models.Goal, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Goal) bool {
		return x.Name == name
	}).([]*models.Goal)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *goalCachedService) FindByID(id int64) (*models.Goal, error) {
	get, err := s.db.Get(GoalsBucket, strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}
	g := &models.Goal{}
	_ = json.Unmarshal(get, &g)
	return g, nil
}

func (s *goalCachedService) Clean() error {
	err := s.db.Truncate(GoalsBucket)
	if err != nil {
		return err
	}
	return nil
}
