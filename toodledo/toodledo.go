package toodledo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultBaseURL = "https://api.toodledo.com"
)

type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	BaseURL     *url.URL
	accessToken string

	// TODO

	common Service

	FolderService *FolderService
	TaskService   *TaskService
	// TODO
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req.WithContext(ctx)
	query := req.URL.Query()
	// toodledo api document: https://api.toodledo.com/3/account/index.php
	query.Add("access_token", c.accessToken)
	req.URL.RawQuery = query.Encode()
	
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
		
	}
    defer resp.Body.Close()
	
	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}
	response := &Response{resp}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return response, err
}

type ApiError struct {
	Response *http.Response
	Body string
	Message  string         `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%v %d %v %v",
		e.Response.Request.Method, e.Response.StatusCode, e.Message, e.Body)
}

func CheckResponse(r * http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	data, _ := ioutil.ReadAll(r.Body)
	return &ApiError{Response: r, Body: string(data), Message: "error"}
}

type Service struct {
	client *Client
}


func NewClient(accessToken string) *Client {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{client: httpClient, BaseURL: baseURL, accessToken: accessToken}
	client.common.client = client // TODO why
	client.FolderService = (*FolderService)(&client.common)
	client.TaskService = (*TaskService)(&client.common)
	// TODO
	return client
}

type Response struct {
	*http.Response
	// TODO
}
