//go:build integration
// +build integration

package suites

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_contextService_Get(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()

	contexts, _, body, err := client.ContextService.Get()
	assert.NoError(t, err)
	assert.Greater(t, len(contexts), 0, fmt.Sprintf("resp: %s", body))
}

func Test_contextService_Add(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()

	context, _, body, err := client.ContextService.Add("test-123")
	assert.NoError(t, err)
	assert.NotNil(t, context, fmt.Sprintf("resp: %s", body))
}

func Test_contextService_Edit(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()

	context, _, body, err := client.ContextService.Edit(1292681, "test-5")
	assert.NoError(t, err)
	assert.NotNil(t, context, fmt.Sprintf("resp: %s", body))
}

func Test_contextService_Delete(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	client := ClientForTest()

	context, body, err := client.ContextService.Delete(1297923)
	assert.NoError(t, err)
	assert.NotNil(t, context, fmt.Sprintf("resp: %s", body))
}
