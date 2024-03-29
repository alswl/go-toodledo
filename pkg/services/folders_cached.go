package services

import (
	"encoding/json"
	"sync"

	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/thoas/go-funk"
)

var FolderBucket = "folders"

type folderCachedService struct {
	syncLock sync.Mutex

	svc        FolderService
	db         dal.Backend
	accountSvc AccountExtService
}

var folderLocalServiceInstance FolderPersistenceService
var folderLocalServiceOnce sync.Once

func NewFolderCachedService(folderSvc FolderService, accountSvc AccountExtService,
	db dal.Backend) FolderPersistenceService {
	s := folderCachedService{
		svc:        folderSvc,
		db:         db,
		accountSvc: accountSvc,
	}
	return &s
}

func ProvideFolderCachedService(
	svc FolderService,
	accountSvc AccountExtService,
	db dal.Backend,
) FolderPersistenceService {
	if folderLocalServiceInstance == nil {
		folderLocalServiceOnce.Do(func() {
			folderLocalServiceInstance = NewFolderCachedService(svc, accountSvc, db)
		})
	}
	return folderLocalServiceInstance
}

func (s *folderCachedService) Sync() error {
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
		_ = s.db.Put(FolderBucket, f.Name, bytes)
	}
	return nil
}

func (s *folderCachedService) PartialSync(lastEditTime *int32) error {
	return s.Sync()
}

func (s *folderCachedService) Rename(name string, newName string) (*models.Folder, error) {
	_ = s.Clean()
	return s.svc.Rename(name, newName)
}

func (s *folderCachedService) Archive(id int, isArchived bool) (*models.Folder, error) {
	_ = s.Clean()
	return s.svc.Archive(id, isArchived)
}

func (s *folderCachedService) Delete(name string) error {
	_ = s.Clean()
	return s.svc.Delete(name)
}

func (s *folderCachedService) Create(name string) (*models.Folder, error) {
	_ = s.Clean()
	return s.svc.Create(name)
}

func (s *folderCachedService) ListAll() ([]*models.Folder, error) {
	fs := make([]*models.Folder, 0)
	all, err := s.db.List(FolderBucket)
	if err != nil {
		return nil, err
	}
	for _, bytes := range all {
		f := &models.Folder{}
		_ = json.Unmarshal(bytes, &f)
		fs = append(fs, f)
	}
	return fs, nil
}

func (s *folderCachedService) Find(name string) (*models.Folder, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Folder) bool {
		return x.Name == name
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *folderCachedService) FindByID(id int64) (*models.Folder, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Folder) bool {
		return x.ID == id
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *folderCachedService) Clean() error {
	err := s.db.Truncate(FolderBucket)
	if err != nil {
		return err
	}
	return nil
}
