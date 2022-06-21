package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/kubectl/pkg/cmd/util/editor"
	"os"
	"strconv"
)

var editorCmd = &cobra.Command{
	Use:  "editor",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// services
		_, err := injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		taskSvc, _ := injector.InitTaskService()
		contextSvc, _ := injector.InitContextCachedService()
		folderSvc, _ := injector.InitFolderCachedService()
		goalSvc, _ := injector.InitGoalCachedService()
		taskRichSvc := services.NewTaskRichService(taskSvc, folderSvc, contextSvc, goalSvc, logging.ProvideLogger())

		// fetch task
		id, _ := strconv.Atoi(args[0])
		t, err := taskSvc.FindById(int64(id))
		if err != nil {
			logrus.WithError(err).Fatal("find task")
			return
		}

		e := editor.NewDefaultEditor([]string{})
		tmpFile := fmt.Sprintf("/tmp/toodledo-task-editor-%d.yaml", t.ID)
		f, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, 0755)
		bs, _ := yaml.Marshal(t)
		f.Write(bs)
		f.Close()

		err = e.Launch(tmpFile)
		if err != nil {
			logrus.WithError(err).Fatal("launch editor")
			return
		}
		f, err = os.OpenFile(tmpFile, os.O_RDONLY, 0644)
		defer f.Close()
		if err != nil {
			logrus.WithError(err).Fatal("open file")
			return
		}
		_, err = f.Read(bs)
		if err != nil {
			logrus.WithError(err).Fatal("read file")
			return
		}
		var inputT models.Task
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
		fmt.Println(render.Tables4RichTasks([]*models.RichTask{rt}))
	},
}

func init() {
	(&cobra.Command{
		Use:   "task",
		Short: "Manage toodledo tasks",
	}).AddCommand(editorCmd)
}
