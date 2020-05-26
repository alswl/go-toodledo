package toodledo

import (
	"context"
)

type FolderService Service

type Folder struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Private  int    `json:"private"`
	Archived int    `json:"archived"`
	Ord      int    `json:"ord"`
}

func (s *FolderService) Get(ctx context.Context) ([]*Folder, *Response, error) {
	path := "/3/folders/get.php"
	
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var folders []*Folder
	resp, err := s.client.Do(ctx, req, &folders)
	if err != nil {
		return nil, resp, err
	}

	return folders, resp, nil
}
