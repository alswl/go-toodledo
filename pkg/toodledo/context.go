package toodledo

import (
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"net/http"
)

// TODO @alswl impl
type ContextService interface {
	Get() ([]models.Context, *http.Response, string, error)
	Add(name string) (models.Context, *http.Response, string, error)
	AddExtend(name string, private bool) (models.Context, *http.Response, string, error)
	Edit(id int, name string, private bool) (models.Context, *http.Response, string, error)
	Delete(id int) (*http.Response, string, error)
}

type contextService Service

func (c *contextService) Get() ([]models.Context, *http.Response, string, error) {
	path := "/3/contexts/get.php"
	var contexts []models.Context
	resp, body, errs := c.client.requests.Get(DefaultBaseUrl + path).Auth(c.client.accessToken).EndStructWithTError(&contexts)
	if len(errs) > 0 {
		firstErr := errs[0]
		return nil, resp, string(body), firstErr
	}
	return contexts, resp, string(body), nil
}

func (c *contextService) AddExtend(name string, private bool) (models.Context, *http.Response, string, error) {
	path := "/3/contexts/add.php"
	var contexts []models.Context
	req := c.client.requests.Post(DefaultBaseUrl + path).Auth(c.client.accessToken)
	req.Send("name=" + name)
	if private {
		req.Send("private=" + "1")
	}

	resp, body, errs := req.EndStructWithTError(&contexts)

	if len(errs) > 0 || len(contexts) == 0 {
		firstErr := errs[0]
		return models.Context{}, resp, string(body), firstErr
	}
	return contexts[0], resp, string(body), nil
}

func (c *contextService) Add(name string) (models.Context, *http.Response, string, error) {
	return c.AddExtend(name, false)
}

func (c *contextService) Edit(id int, name string, private bool) (models.Context, *http.Response, string, error) {
	panic("implement me")
}

func (c *contextService) Delete(id int) (*http.Response, string, error) {
	panic("implement me")
}
