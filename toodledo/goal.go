package toodledo

import (
	"context"
	"net/url"
	"strconv"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

type GoalService Service

type GoalLevel int

const (
	GOAL_LEVEL_LIFE_TIME  int = 0
	GOAL_LEVEL_LONG_TERM  int = 1
	GOAL_LEVEL_SHORT_TERM int = 2
)

type Goal struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Level GoalLevel `json:"level"`
	// 0 or 1
	Archived    int    `json:"archived"`
	Contributes int    `json:"contributes"`
	Note        string `json:"note"`
}

var validate *validator.Validate

type GoalAdd struct {
	// required
	Name  string `validate:"required`
	Level *GoalLevel
	// 0 or 1
	Contributes *int
	// 0 or 1
	Private *bool
	Note    *string
}

func (s *GoalService) Get(ctx context.Context) ([]*Goal, *Response, error) {
	path := "/3/goals/get.php"

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	var goals []*Goal
	resp, err := s.client.Do(ctx, req, &goals)
	if err != nil {
		return nil, resp, err
	}

	return goals, resp, nil
}

func (s *GoalService) Add(ctx context.Context, goalAdd GoalAdd) (*Goal, *Response, error) {
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

	var goals []*Goal
	resp, err := s.client.Do(ctx, req, &goals)
	log.Warn(resp, err)
	
	if err != nil {
		log.WithFields(log.Fields{"resp": resp, "err": err}).Warn("err")
		return nil, resp, err
	}

	// get first
	return goals[0], resp, nil
}
