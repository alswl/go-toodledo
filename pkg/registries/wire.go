package registries

import (
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(auth.ProvideSimpleAuth, auth.ProvideAccessToken, auth.ProvideOAuth2Config())
