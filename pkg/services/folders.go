package services

import (
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type FolderService interface {
	FindByName(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
}

type folderService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewFolderService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) FolderService {
	return &folderService{cli: cli, auth: auth}
}

func NewFolderService0(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) *folderService {
	return &folderService{cli: cli, auth: auth}
}

func (s *folderService) FindByName(name string) (*models.Folder, error) {
	panic("implement me")
}

func (s *folderService) ListAll() ([]*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	return ts.Payload, nil
}
