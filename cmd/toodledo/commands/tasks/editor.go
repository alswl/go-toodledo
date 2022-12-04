package tasks

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/kubectl/pkg/cmd/util/editor"
)

func NewEditorCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "editor",
		Args:  cobra.ExactArgs(1),
		Short: "Edit a task in editor",
		Example: heredoc.Doc(`
		$ toodledo tasks editor 8848
	`),
		Run: func(cmd *cobra.Command, args []string) {
			// services
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			taskSvc := app.TaskSvc
			// contextSvc := app.ContextSvc
			// folderSvc := app.FolderSvc
			// goalSvc := app.GoalSvc
			taskRichSvc := app.TaskRichSvc

			// fetch task
			id, _ := strconv.Atoi(args[0])
			t, err := taskSvc.FindByID(int64(id))
			if err != nil {
				logrus.WithError(err).Fatal("find task")
				return
			}

			e := editor.NewDefaultEditor([]string{})
			tmpFilePath := fmt.Sprintf("/tmp/toodledo-task-editor-%d.yaml", t.ID)
			tmpFile, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_RDWR, 0755)
			if err != nil {
				logrus.WithError(err).Fatal("open tmp file")
				return
			}
			bs, _ := yaml.Marshal(t)
			_, err = tmpFile.Write(bs)
			if err != nil {
				logrus.WithError(err).Fatal("write task to tmp file")
				return
			}
			err = tmpFile.Close()
			if err != nil {
				logrus.WithError(err).Fatal("close tmp file")
			}

			err = e.Launch(tmpFilePath)
			if err != nil {
				logrus.WithError(err).Fatal("launch editor")
				return
			}
			tmpFile, err = os.OpenFile(tmpFilePath, os.O_RDONLY, 0644)
			defer func() {
				err = tmpFile.Close()
				if err != nil {
					logrus.WithError(err).Fatal("close tmp file")
				}
			}()
			if err != nil {
				logrus.WithError(err).Fatal("open file")
				return
			}
			_, err = tmpFile.Read(bs)
			if err != nil {
				logrus.WithError(err).Fatal("read file")
				return
			}
			var inputT models.TaskEdit
			err = yaml.Unmarshal(bs, &inputT)
			if err != nil {
				logrus.WithError(err).Fatal("unmarshal yaml")
				return
			}

			newT, err := taskSvc.Edit(int64(id), &inputT)
			if err != nil {
				logrus.WithError(err).Fatal("edit task")
			}

			rt, _ := taskRichSvc.Rich(newT)
			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4RichTasks([]*models.RichTask{rt}))
		},
	}
}
