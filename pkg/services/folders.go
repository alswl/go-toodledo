package services

import (
	"errors"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/thoas/go-funk"
	"strconv"
)

type FolderService interface {
	FindByName(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
	ArchiveFolder(id int, isArchived bool) (*models.Folder, error)
}

type folderService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewFolderService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) *folderService {
	return &folderService{cli: cli, auth: auth}
}

func (s *folderService) FindByName(name string) (*models.Folder, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Folder) bool {
		return x.Name == name
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}

func (s *folderService) ListAll() ([]*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	return ts.Payload, nil
}

func (s *folderService) ArchiveFolder(id int, isArchived bool) (*models.Folder, error) {
	// TODO test
	cli := client.NewHTTPClient(strfmt.NewFormats())
	p := folder.NewPostFoldersEditPhpParams()
	p.SetID(strconv.Itoa(id))
	archived := int64(0)
	if isArchived {
		archived = 1
	}
	p.SetArchived(&archived)
	res, err := cli.Folder.PostFoldersEditPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload[0], nil
}
