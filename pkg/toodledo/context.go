package toodledo

import (
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"net/http"
	"strconv"
)

type ContextService interface {
	// Get return all contexts
	Get() ([]models.Context, *http.Response, string, error)

	// Add a new context, name is context name, default context is public
	Add(name string) (models.Context, *http.Response, string, error)

	// Add a new context, with parameter name, private
	AddExtend(name string, private bool) (models.Context, *http.Response, string, error)

	// Edit a context name by id
	Edit(id int, name string) (models.Context, *http.Response, string, error)

	// EditExtend a context name and private by id
	EditExtend(id int, name string, private bool) (models.Context, *http.Response, string, error)

	// Delete a context
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

func (c *contextService) Edit(id int, name string) (models.Context, *http.Response, string, error) {
	return c.EditExtend(id, name, false)
}

func (c *contextService) EditExtend(id int, name string, private bool) (models.Context, *http.Response, string, error) {
	path := "/3/contexts/edit.php"
	var contexts []models.Context
	req := c.client.requests.Post(DefaultBaseUrl + path).Auth(c.client.accessToken)
	req.Send("id=" + strconv.Itoa(id))
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

func (c *contextService) Delete(id int) (*http.Response, string, error) {
	path := "/3/contexts/delete.php"
	var contexts []models.Context
	req := c.client.requests.Post(DefaultBaseUrl + path).Auth(c.client.accessToken)
	req.Send("id=" + strconv.Itoa(id))

	resp, body, errs := req.EndStructWithTError(&contexts)

	if len(errs) > 0 || len(contexts) == 0 {
		firstErr := errs[0]
		return resp, string(body), firstErr
	}
	return resp, string(body), nil
}
