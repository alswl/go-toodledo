package services

import (
	"errors"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/thoas/go-funk"
	"strconv"
)

// TODO using service

func FindGoalByName(auth runtime.ClientAuthInfoWriter, name string) (*models.Goal, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), auth)
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

func ArchiveGoal(auth runtime.ClientAuthInfoWriter, id int, isArchived bool) (*models.Goal, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	p := goal.NewPostGoalsEditPhpParams()
	p.SetID(strconv.Itoa(id))
	archived := int64(0)
	if isArchived {
		archived = 1
	}
	p.SetArchived(&archived)
	res, err := cli.Goal.PostGoalsEditPhp(p, auth)
	if err != nil {
		return nil, err
	}
	return res.Payload[0], nil
}
