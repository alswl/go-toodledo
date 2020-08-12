package integration

import (
	"context"
	toodledo "github.com/alswl/go-toodledo/pkg"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoalService_Get(t *testing.T) {
	accessToken := "***REMOVED***"
	assert.NotNil(t, accessToken)

	client := toodledo.NewClient(accessToken)
	ctx := context.Background()
	elems, _, err := client.GoalService.Get(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, elems)
	first := elems[0]
	assert.Equal(t, first.Name, "goal-abc")
}

func TestGoalService_Add(t *testing.T) {
	accessToken := "***REMOVED***"
	assert.NotNil(t, accessToken)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	goal := toodledo.GoalAdd{Name: "goal-b"}
	client := toodledo.NewClient(accessToken)
	ctx := context.Background()
	elem, _, err := client.GoalService.Add(ctx, goal)
	assert.Nil(t, err)
	assert.NotNil(t, elem)
}
