package toodledo

import "context"

type FolderService service

type Folder struct {
}

func (s *FolderService) Get(ctx context.Context) (*Folder, *Response, error) {
	req, err := s.client.NewRequest("GET", "folder", nil)
	if err != nil {
		return nil, nil, err
	}

	f := &Folder{}
	resp, err := s.client.Do(ctx, req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
