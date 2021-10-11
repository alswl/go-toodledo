package services

import (
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

func (s *folderService) ListAll() ([]*models.Folder, error) {
	// XXX using bolt
	return s.listAll()
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
