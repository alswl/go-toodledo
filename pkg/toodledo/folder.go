package toodledo

import (
	"context"
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"net/url"
	"strconv"
)

// FolderService ...
type FolderService Service

// Get ...
func (s *FolderService) Get(ctx context.Context) ([]*models.Folder, *Response, error) {
	path := "/3/folders/get.php"

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	var folders []*models.Folder
	resp, err := s.client.Do(ctx, req, &folders)
	if err != nil {
		return nil, resp, err
	}

	return folders, resp, nil
}

// Add ...
func (s *FolderService) Add(ctx context.Context, name string) (*models.Folder, *Response, error) {
	path := "/3/folders/add.php"

	form := url.Values{}
	form.Add("name", name)
	req, err := s.client.NewRequestWithForm("POST", path, form)
	if err != nil {
		return nil, nil, err
	}

	var folders []*models.Folder
	resp, err := s.client.Do(ctx, req, &folders)
	if err != nil {
		return nil, resp, err
	}

	// return first folder, this API always return one folder.
	return folders[0], resp, nil
}

// Edit ...
func (s *FolderService) Edit(ctx context.Context, id int, name string) (*models.Folder, *Response, error) {
	return s.EditWithPrivate(ctx, id, name, -1)

}

// EditWithPrivate ...
func (s *FolderService) EditWithPrivate(ctx context.Context, id int, name string, private int) (*models.Folder, *Response, error) {
	path := "/3/folders/edit.php"

	form := url.Values{}
	form.Add("id", strconv.Itoa(id))
	form.Add("name", name)
	if private != -1 {
		form.Add("private", fmt.Sprint(private))
	}

	req, err := s.client.NewRequestWithForm("POST", path, form)
	if err != nil {
		return nil, nil, err
	}

	var folders []*models.Folder
	resp, err := s.client.Do(ctx, req, &folders)
	if err != nil {
		return nil, resp, err
	}

	// return first folder, this API always return one folder.
	return folders[0], resp, nil
}

// Delete ...
func (s *FolderService) Delete(ctx context.Context, id int) (*Response, error) {
	path := "/3/folders/delete.php"

	form := url.Values{}
	form.Add("id", strconv.Itoa(id))

	req, err := s.client.NewRequestWithForm("POST", path, form)
	if err != nil {
		return nil, err
	}

	var result *map[string]int
	resp, err := s.client.Do(ctx, req, &result)
	if err != nil {
		return resp, err
	}
	if (*result)["deleted"] != id {
		return resp, errors.New("delete failed")
	}

	// return first folder, this API always return one folder.
	return resp, nil
}
