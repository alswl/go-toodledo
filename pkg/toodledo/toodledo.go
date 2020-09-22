package toodledo

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	FolderService  *FolderService
	GoalService    *GoalService
	TaskService    *TaskService
	// TODO
}

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, map[string]string{}, url.Values{})
}

func (c *Client) NewRequestWithParams(method, urlStr string, params map[string]string) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, params, url.Values{})
}

func (c *Client) NewRequestWithForm(method, urlStr string, form url.Values) (*http.Request, error) {
	return c.NewRequestWithParamsAndForm(method, urlStr, map[string]string{}, form)
}

func (c *Client) NewRequestWithParamsAndForm(method, urlStr string, params map[string]string, form url.Values) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if len(params) != 0 {
		q := u.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
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

// Do request to toodledo API server and unmarshalled, req is request, v is response unmarshalled, Response is http response
// toodledo api document: https://api.toodledo.com/3/account/index.php
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req.WithContext(ctx)
	query := req.URL.Query()
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

	err = CheckResponseStatus(resp)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	if err != nil {
		return nil, err
	}
	if toodledoError, ok := CheckToodledoResponse(body); !ok {
		return &Response{
				Response: resp,
				Text:     toodledoError.ErrorDesc,
			}, ApiError{
				Response: nil,
				Body:     body,
				Message:  toodledoError.ErrorDesc,
			}
	}
	response := &Response{resp, body}
	log.WithFields(log.Fields{"response": response.Text}).Debug("requested toodledo")
	if v != nil {
		if writer, ok := v.(io.Writer); ok {
			io.Copy(writer, strings.NewReader(body))
		} else {
			decErr := json.NewDecoder(strings.NewReader(body)).Decode(v)
			if decErr != nil {
				log.Warn("decErr: ", decErr)
				log.Warn("v: ", v)
			}
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				return response, decErr
			}
		}
	}
	return response, nil
}

type ApiError struct {
	Response *http.Response
	Body     string
	Message  string `json:"message"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%v %d %v %v",
		e.Response.Request.Method, e.Response.StatusCode, e.Message, e.Body)
}

func CheckResponseStatus(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return &ApiError{Response: r, Body: string(data), Message: "http status error"}
}

// CheckToodledoResponse check body is Toodledo Error Response format
func CheckToodledoResponse(body string) (ErrorResponse, bool) {
	errorResponse := ErrorResponse{}
	err := json.NewDecoder(strings.NewReader(body)).Decode(&errorResponse)
	if err == nil && errorResponse.ErrorCode != 0 {
		return errorResponse, false
	}
	return ErrorResponse{}, true
}

type Service struct {
	client *Client
}

func NewClient(accessToken string) *Client {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{client: httpClient, BaseURL: baseURL, accessToken: accessToken}
	client.common.client = client
	client.AccountService = (*AccountService)(&client.common)
	client.FolderService = (*FolderService)(&client.common)

	var ts TaskService = &taskService{client}
	client.TaskService = &ts
	client.GoalService = (*GoalService)(&client.common)
	// TODO
	return client
}

type Response struct {
	*http.Response
	// TODO
	Text string
}
