//go:build integration

package suites

import (
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_contextService_Get(t *testing.T) {
	app, err := itinjector.InitCLIApp()
	assert.NoError(t, err)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc := app.ContextSvc

	contexts, err := svc.ListAll()
	assert.NoError(t, err)
	assert.Greater(t, len(contexts), 0)
}

func Test_contextService_Add(t *testing.T) {
	app, err := itinjector.InitCLIApp()
	assert.NoError(t, err)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc := app.ContextSvc

	context, err := svc.Create("test-123")
	assert.NoError(t, err)
	assert.NotNil(t, context)
}

func Test_contextService_Edit(t *testing.T) {
	app, err := itinjector.InitCLIApp()
	assert.NoError(t, err)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc := app.ContextSvc

	context, err := svc.Rename("test-123", "test-5")
	assert.NoError(t, err)
	assert.NotNil(t, context)
}

func Test_contextService_Delete(t *testing.T) {
	app, err := itinjector.InitCLIApp()
	assert.NoError(t, err)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	svc := app.ContextSvc

	err = svc.Delete("test-5")
	assert.NoError(t, err)
}
