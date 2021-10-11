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

type FolderService interface {
	FindByName(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
}

type folderService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
	db   *bolt.DB
}

func NewFolderService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter, db *bolt.DB) FolderService {
	return &folderService{cli: cli, auth: auth, db: db}
}

func (s *folderService) listAll() ([]*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	return ts.Payload, nil
}

func (s *folderService) listAllFromDB() ([]*models.Folder, error) {
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

func (s *folderService) put2DB(folders []*models.Folder) error {
	s.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("folders"))
		for _, f := range folders {
			bytes, _ := json.Marshal(f)
			b.Put(([]byte)(strconv.Itoa((int)(f.ID))), bytes)
		}
		return nil
	})
	return nil
}

func (s *folderService) ListAll() ([]*models.Folder, error) {
	// XXX using bolt
	all, err := s.listAll()
	if err != nil {
		return nil, err
	}
	s.put2DB(all)
	return s.listAllFromDB()
}

func (s *folderService) Find(id int) (*models.Folder, error) {
	// TODO
	return nil, nil
}

func (s *folderService) FindByName(name string) (*models.Folder, error) {
	// TODO
	return nil, nil
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
