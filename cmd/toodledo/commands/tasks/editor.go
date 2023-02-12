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
	"github.com/alswl/go-toodledo/pkg/utils/editor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

			e, err := editor.NewDefaultEditor()
			if err != nil {
				logrus.WithError(err).Fatal("find EDITOR")
				return
			}
			tmpFilePath := fmt.Sprintf("/tmp/toodledo-task-editor-%d.yaml", t.ID)
			e.CleanScience(tmpFilePath)
			// clean tmpFile
			defer func() {
				e.CleanScience(tmpFilePath)
			}()
			tmpFile, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_RDWR, 0755)
			if err != nil {
				logrus.WithError(err).Fatal("open tmp file")
				return
			}
			bs := models.PrettyYAML(t)
			_, err = tmpFile.Write([]byte(bs))
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
			var newBs []byte
			newBs, err = os.ReadFile(tmpFilePath)
			if err != nil {
				logrus.WithError(err).Fatal("read file")
				return
			}
			inputT, err := models.LoadTaskFromYAML(string(newBs))
			if err != nil {
				logrus.WithError(err).Fatal("unmarshal yaml")
				return
			}

			newT, err := taskSvc.Edit(int64(id), inputT)
			if err != nil {
				logrus.WithError(err).Fatal("edit task")
			}

			rt, _ := taskRichSvc.Rich(newT)
			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4RichTasks([]*models.RichTask{rt}))
		},
	}
}
