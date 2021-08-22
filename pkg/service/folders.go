package service

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
}

type FolderServiceImpl struct {
	client *client.Toodledo
}

func NewFolderServiceImpl(client *client.Toodledo) *FolderServiceImpl {
	return &FolderServiceImpl{client: client}
}

func (s *FolderServiceImpl) Find(id int) (*models.Folder, error) {
	// TODO
	return nil, nil
}

func (s *FolderServiceImpl) FindByName(name string) (*models.Folder, error) {

	// TODO
	return nil, nil
}

func FindFolderByName(auth runtime.ClientAuthInfoWriter, name string) (*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
	if err != nil {
		return nil, err
	}
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
