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

// FolderService ...
type FolderService interface {
	Find(name string) (*models.Folder, error)
	ListAll() ([]*models.Folder, error)
	Rename(name string, newName string) (*models.Folder, error)
	Archive(id int, isArchived bool) (*models.Folder, error)
	Delete(name string) error
	Create(name string) (*models.Folder, error)
}

// folderService query folder with remote api
type folderService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

// NewFolderService ...
func NewFolderService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) FolderService {
	return &folderService{cli: cli, auth: auth}
}

// Create ...
func (s *folderService) Create(name string) (*models.Folder, error) {
	params := folder.NewPostFoldersAddPhpParams()
	params.SetName(name)
	resp, err := s.cli.Folder.PostFoldersAddPhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("Failed to create")
		return nil, err
	}
	return resp.Payload[0], nil
}

// Delete ...
func (s *folderService) Delete(name string) error {
	f, err := s.Find(name)
	if err != nil {
		return err
	}

	params := folder.NewPostFoldersDeletePhpParams()
	params.SetID(f.ID)
	resp, err := s.cli.Folder.PostFoldersDeletePhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("Failed to delete folder")
		return err
	}
	return nil
}

// Rename ...
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

// Find ...
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

// ListAll ...
func (s *folderService) ListAll() ([]*models.Folder, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	return ts.Payload, nil
}

// Archive ...
func (s *folderService) Archive(id int, isArchived bool) (*models.Folder, error) {
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
