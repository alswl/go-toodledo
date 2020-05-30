package integration

import (
	"context"
	"github.com/alswl/go-toodledo/toodledo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountService_Get(t *testing.T) {
	accessToken := "***REMOVED***"
	assert.NotNil(t, accessToken)

	client := toodledo.NewClient(accessToken)
	ctx := context.Background()
	account, _, err := client.AccountService.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, account.Userid, "***REMOVED***")
	assert.Equal(t, account.Alias, "***REMOVED***")
	assert.Equal(t, account.Email, "***REMOVED***")
}
