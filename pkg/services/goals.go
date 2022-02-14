package services

import (
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
)

// GoalService ...
type GoalService interface {
	Find(name string) (*models.Goal, error)
	FindByID(id int64) (*models.Goal, error)
	Archive(id int, isArchived bool) (*models.Goal, error)
	Delete(name string) error
	Rename(name string, newName string) (*models.Goal, error)
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

// Find ...
func (s *goalService) Find(name string) (*models.Goal, error) {
	logrus.Warn("FindByID is implemented with ListALl(), it's deprecated, please using cache")
	ts, err := s.cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), s.auth)
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

func (s *goalService) FindByID(id int64) (*models.Goal, error) {
	logrus.Warn("FindByID is implemented with ListALl(), it's deprecated, please using cache")
	ts, err := s.cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	filtered := funk.Filter(ts.Payload, func(x *models.Goal) bool {
		return x.ID == id
	}).([]*models.Goal)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}

// Archive ...
func (s *goalService) Archive(id int, isArchived bool) (*models.Goal, error) {
	p := goal.NewPostGoalsEditPhpParams()
	p.SetID(strconv.Itoa(id))
	archived := int64(0)
	if isArchived {
		archived = 1
	}
	p.SetArchived(&archived)
	res, err := s.cli.Goal.PostGoalsEditPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload[0], nil
}

// Create ...
func (s *goalService) Create(name string) (*models.Goal, error) {
	params := goal.NewPostGoalsAddPhpParams()
	params.SetName(name)
	res, err := s.cli.Goal.PostGoalsAddPhp(params, s.auth)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return res.Payload[0], nil
}

// Delete ...
func (s *goalService) Delete(name string) error {
	g, err := s.Find(name)
	if err != nil {
		return err
	}

	params := goal.NewPostGoalsDeletePhpParams()
	params.SetID(g.ID)
	resp, err := s.cli.Goal.PostGoalsDeletePhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).Error(err)
		return err
	}
	return nil
}

// Rename ...
func (s *goalService) Rename(name string, newName string) (*models.Goal, error) {
	if name == newName {
		logrus.Error("not changed")
		return nil, fmt.Errorf("not changed")
	}
	g, err := s.Find(name)
	if err != nil {
		logrus.Error(err)
		return nil, common.ErrNotFound
	}

	p := goal.NewPostGoalsEditPhpParams()
	p.SetID(strconv.Itoa(int(g.ID)))
	p.SetName(&newName)
	res, err := s.cli.Goal.PostGoalsEditPhp(p, s.auth)
	if err != nil {
		logrus.WithField("resp", res).WithError(err).Error("request failed")
		return nil, err
	}
	return res.Payload[0], nil
}

// ListAll ...
func (s *goalService) ListAll() ([]*models.Goal, error) {
	res, err := s.cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), s.auth)
	if err != nil {
		logrus.WithField("resp", res).WithError(err).Error("request failed")
		return nil, err
	}
	return res.Payload, nil
}
