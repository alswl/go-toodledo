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
	"strings"
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

	AccountService *AccountService
	FolderService *FolderService
	TaskService   *TaskService
	// TODO
}

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, map[string]string{}, url.Values{})
}

func (c *Client) NewRequestWithParams(method, urlStr string, params map[string]string) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, params, url.Values{})
}

func (c *Client) NewRequestWithForm(method, urlStr string, form url.Values) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, map[string]string{}, url.Values{})
}

func (c *Client) NewRequestWithParamsAndForm(method, urlStr string, params map[string]string, form url.Values) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if len(params) != 0 {
		for k, v := range params {
			u.Query().Add(k, v)
		}
		u.RawQuery = u.Query().Encode()
	}

	var buf io.Reader
	if len(form) != 0 {
		buf = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if form != nil {
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response := &Response{resp, string(bodyBytes)}
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
	Body     string
	Message  string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%v %d %v %v",
		e.Response.Request.Method, e.Response.StatusCode, e.Message, e.Body)
}

func CheckResponse(r *http.Response) error {
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
	client.AccountService = (*AccountService)(&client.common)
	client.FolderService = (*FolderService)(&client.common)
	client.TaskService = (*TaskService)(&client.common)
	// TODO
	return client
}

type Response struct {
	*http.Response
	// TODO
	Text string
}
