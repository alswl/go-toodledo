package cmdutil

import (
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/iostreams"
	"github.com/alswl/go-toodledo/pkg/models"
	utilsos "github.com/alswl/go-toodledo/pkg/utils/os"
)

type Browser interface {
	Browse(string) error
}

type DefaultBrowser struct {
}

func (b DefaultBrowser) Browse(s string) error {
	return utilsos.OpenInBrowser(s)
}

type Factory struct {
	IOStreams *iostreams.IOStreams
	Browser   Browser
	Config    func() (models.ToodledoCliConfig, error)

	ExecutableName string
}

func NewFactory() *Factory {
	return &Factory{
		IOStreams:      iostreams.UsingSystem(),
		Browser:        DefaultBrowser{},
		Config:         common.NewCliConfigFromViper,
		ExecutableName: "toodledo",
	}

}
