package services

import (
	"encoding/json"
	"errors"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

type folderCachedService struct {
	svc *folderService
	db  *bolt.DB
}

func NewFolderCachedService(svc0 *folderService, db *bolt.DB) FolderService {
	return &folderCachedService{svc: svc0, db: db}
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

func (s folderCachedService) syncIfExpired() error {
	// TODO if expired
	all, err := s.svc.ListAll()
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

func FindFolderByName(auth runtime.ClientAuthInfoWriter, name string) (*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
	if err != nil {
		return nil, err
	}
	// TODO replace with svc.ListAll()

	filtered := funk.Filter(ts.Payload, func(x *models.Folder) bool {
		return x.Name == name
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}

func ArchiveFolder(auth runtime.ClientAuthInfoWriter, id int, isArchived bool) (*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	p := folder.NewPostFoldersEditPhpParams()
	p.SetID(strconv.Itoa(id))
	archived := int64(0)
	if isArchived {
		archived = 1
	}
	p.SetArchived(&archived)
	res, err := cli.Folder.PostFoldersEditPhp(p, auth)
	if err != nil {
		return nil, err
	}
	return res.Payload[0], nil
}
