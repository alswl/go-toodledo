package services

import (
	"errors"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
)

// GoalService ...
type GoalService interface {
	FindByName(name string) (*models.Goal, error)
	Archive(id int, isArchived bool) (*models.Goal, error)
	Delete(id int64) error
	Rename(id int64, newName string) (*models.Goal, error)
	Create(name string) (*models.Goal, error)
	ListAll() ([]*models.Goal, error)
}

type goalService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
	//log  *logrus.Logger
}

// NewGoalService ...
func NewGoalService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) GoalService {
	return &goalService{cli: cli, auth: auth}
}

// FindByName ...
func (h *goalService) FindByName(name string) (*models.Goal, error) {
	ts, err := h.cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), h.auth)
	if err != nil {
		return nil, err
	}
	filtered := funk.Filter(ts.Payload, func(x *models.Goal) bool {
		return x.Name == name
	}).([]*models.Goal)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}

// Archive ...
func (h *goalService) Archive(id int, isArchived bool) (*models.Goal, error) {
	p := goal.NewPostGoalsEditPhpParams()
	p.SetID(strconv.Itoa(id))
	archived := int64(0)
	if isArchived {
		archived = 1
	}
	p.SetArchived(&archived)
	res, err := h.cli.Goal.PostGoalsEditPhp(p, h.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload[0], nil
}

// Create ...
func (h *goalService) Create(name string) (*models.Goal, error) {
	params := goal.NewPostGoalsAddPhpParams()
	params.SetName(name)
	res, err := h.cli.Goal.PostGoalsAddPhp(params, h.auth)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return res.Payload[0], nil
}

// Delete ...
func (h *goalService) Delete(id int64) error {
	params := goal.NewPostGoalsDeletePhpParams()
	params.SetID(id)
	resp, err := h.cli.Goal.PostGoalsDeletePhp(params, h.auth)
	if err != nil {
		logrus.WithField("resp", resp).Error(err)
		return err
	}
	return nil
}

// Rename ...
func (h *goalService) Rename(id int64, newName string) (*models.Goal, error) {
	p := goal.NewPostGoalsEditPhpParams()
	p.SetID(strconv.Itoa(int(id)))
	p.SetName(&newName)
	res, err := h.cli.Goal.PostGoalsEditPhp(p, h.auth)
	if err != nil {
		logrus.WithField("resp", res).WithError(err).Error("request failed")
		return nil, err
	}
	return res.Payload[0], nil
}

// ListAll ...
func (h *goalService) ListAll() ([]*models.Goal, error) {
	res, err := h.cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), h.auth)
	if err != nil {
		logrus.WithField("resp", res).WithError(err).Error("request failed")
		return nil, err
	}
	return res.Payload, nil
}
