package client0

import (
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"os"
)

// NewToodledoClient creates a new Toodledo client
// logrus.FieldLogger is forcing in wire to init logging system
func NewToodledoClient(_ logrus.FieldLogger) *Toodledo {
	debug := os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""

	transportConfig := openapiclient.New(DefaultHost, DefaultBasePath, []string{"https"})
	transportConfig.Debug = debug
	// logging to http-request.log
	// FIXME not works now
	transportConfig.SetLogger(logging.GetLoggerOrCreate(constants.LogRequest))
	return New(transportConfig, strfmt.Default)
}
