package folders

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/iostreams"

	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	mockservices "github.com/alswl/go-toodledo/test/mock/services"
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/stretchr/testify/assert"
)

func TestRenameFn(t *testing.T) {
	_, err := itinjector.InitCLIApp()
	if err != nil {
		t.Fatal(err)
	}
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}
	folderSvc := mockservices.NewFolderService(t)
	folderSvc.On("Rename", "reading", "new-name").Return(&models.Folder{Name: "new-name"}, nil)

	renameFn(f, folderSvc, nil, []string{"reading", "new-name"})
	assert.Equal(t, " # │ NAME     │ ARCHIVED \n───┼──────────┼──────────\n 0 │ new-name │        0 \n\n", stdout.String())
}
