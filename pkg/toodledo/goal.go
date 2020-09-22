package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

// TODO @jingchao interface

type GoalService Service

const (
	GOAL_LEVEL_LIFE_TIME  int = 0
	GOAL_LEVEL_LONG_TERM  int = 1
	GOAL_LEVEL_SHORT_TERM int = 2
)

var validate *validator.Validate

func (s *GoalService) Get(ctx context.Context) ([]*models.Goal, *Response, error) {
	path := "/3/goals/get.php"

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	var goals []*models.Goal
	resp, err := s.client.Do(ctx, req, &goals)
	if err != nil {
		return nil, resp, err
	}

	return goals, resp, nil
}

func (s *GoalService) Add(ctx context.Context, goalAdd models.GoalAdd) (*models.Goal, *Response, error) {
	path := "/3/goals/add.php"

	validate = validator.New()
	err := validate.Struct(goalAdd)
	if err != nil {
		return nil, nil, err
	}

	form := url.Values{}
	form.Add("name", goalAdd.Name)
	if goalAdd.Level != nil {
		form.Add("level", strconv.Itoa(int(*goalAdd.Level)))
	}
	if goalAdd.Contributes != nil {
		form.Add("contributes", strconv.Itoa(*goalAdd.Contributes))
	}
	if goalAdd.Private != nil {
		form.Add("private", bool2ints(*goalAdd.Private))
	}
	if goalAdd.Note != nil {
		form.Add("note", *goalAdd.Note)
	}
	req, err := s.client.NewRequestWithForm("POST", path, form)
	if err != nil {
		return nil, nil, err
	}

	var goals []*models.Goal
	resp, err := s.client.Do(ctx, req, &goals)

	if err != nil {
		log.WithFields(log.Fields{"resp": resp, "err": err}).Warn("err")
		return nil, resp, err
	}

	// get first
	return goals[0], resp, nil
}

func (s *GoalService) Edit(ctx context.Context, goalEdit models.GoalEdit) (models.Goal, Response, error) {
	// TODO @alswl
	panic("not impl")
}

func (s *GoalService) Delete(ctx context.Context, id int) (*Response, error) {
	path := "/3/goals/delete.php"

	req, err := s.client.NewRequestWithParams("POST", path, map[string]string{"id": strconv.Itoa(id)})
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		log.WithFields(log.Fields{"resp": resp, "err": err}).Warn("err")
		return resp, err
	}

	return resp, nil
}
