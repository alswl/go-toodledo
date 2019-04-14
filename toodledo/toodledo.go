package toodledo

import (
	"context"
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultBaseURL = "https://api.toodledo.com/3/"
)

type Client struct {
	clientMu sync.Mutex
	client  *http.Client

	BaseURL *url.URL
	// TODO

	common service
	
	Folder *FolderService
	// TODO
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return nil, nil
	// TODO
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	return nil, nil
	// TODO
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{client: httpClient, BaseURL: baseURL}
	client.Folder = (*FolderService)(&client.common)
	// TODO

	return client
}



type Response struct {
	*http.Response
}
