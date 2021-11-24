package services

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
)

type FolderService interface {
	Find(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
	Rename(name string, newName string) (*models.Folder, error)
	ArchiveFolder(id int, isArchived bool) (*models.Folder, error)
}

type folderService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func (s *folderService) Rename(name string, newName string) (*models.Folder, error) {
	if name == newName {
		logrus.Error("not changed")
		return nil, fmt.Errorf("not changed")
	}

	f, err := s.Find(name)
	if err != nil {
		logrus.Error(err)
		return nil, common.ErrNotFound
	}

	p := folder.NewPostFoldersEditPhpParams()
	p.SetID(strconv.Itoa(int(f.ID)))
	p.SetName(&newName)
	resp, err := s.cli.Folder.PostFoldersEditPhp(p, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("request failed")
		return nil, err
	}
	return resp.Payload[0], nil
}

func NewFolderService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) *folderService {
	return &folderService{cli: cli, auth: auth}
}

func (s *folderService) Find(name string) (*models.Folder, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Folder) bool {
		return x.Name == name
	}).([]*models.Folder)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
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
