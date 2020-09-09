//+build integration

package suites

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/toodledo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
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
	
	goal := toodledo.GoalAdd{Name: "goal-b"}
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

	_, err := client.GoalService.Delete(ctx, 510357)
	assert.Nil(t, err)
}
