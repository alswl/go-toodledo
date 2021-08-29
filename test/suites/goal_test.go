//go:build integration
// +build integration

package suites

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoalService_Get(t *testing.T) {
	client := ClientForTest()
	ctx := context.Background()

	elems, _, err := client.GoalService.Get(ctx)

	assert.NoError(t, err)
	assert.NotEmpty(t, elems)
	first := elems[0]
	fmt.Println(first)
	assert.Equal(t, first.Name, "goal-b")
}

func TestGoalService_Add(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()
	ctx := context.Background()

	goal := models.GoalAdd{Name: "goal-b"}
	elem, _, err := client.GoalService.Add(ctx, goal)

	fmt.Println(elem)
	assert.Nil(t, err)
	assert.NotNil(t, elem)
}

func TestGoalService_Delete(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()
	ctx := context.Background()
	now := time.Now()
	nowString := now.Format("20060102150405")

	name := models.GoalAdd{Name: fmt.Sprintf("goal-%s", nowString)}
	// TODO test
	newGoal, _, err := client.GoalService.Add(ctx, name)
	if err != nil {
		panic(err)
	}

	_, err = client.GoalService.Delete(ctx, newGoal.ID)
	assert.Nil(t, err)
}
