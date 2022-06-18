package cmds

import (
	"github.com/alswl/go-toodledo/pkg/iostreams"
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

	ExecutableName string
}

func NewFactory() *Factory {
	return &Factory{
		IOStreams:      iostreams.UsingSystem(),
		Browser:        DefaultBrowser{},
		ExecutableName: "toodledo",
	}

}
