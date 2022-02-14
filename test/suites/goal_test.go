//go:build integration
// +build integration

package suites

import (
	"fmt"
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoalService_Get(t *testing.T) {
	svc, _ := itinjector.InitGoalService()
	all, err := svc.ListAll()

	assert.NoError(t, err)
	assert.NotEmpty(t, all)
	first := all[0]
	fmt.Println(first)
	assert.Equal(t, first.Name, "goal-b")
}

func TestGoalService_Add(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc, _ := itinjector.InitGoalService()

	elem, err := svc.Create("abc")

	fmt.Println(elem)
	assert.Nil(t, err)
	assert.NotNil(t, elem)
}

func TestGoalService_Delete(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc, _ := itinjector.InitGoalService()
	now := time.Now()
	nowString := now.Format("20060102150405")
	// TODO test
	newGoal, err := svc.Create(fmt.Sprintf("goal-%s", nowString))
	if err != nil {
		panic(err)
	}

	err = svc.Delete(newGoal.ID)
	assert.Nil(t, err)
}
