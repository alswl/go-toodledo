package services

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/subtasksview"
	"github.com/thoas/go-funk"
)

func SortSubTasks(tasks []*models.Task, mode subtasksview.Mode) ([]*models.Task, error) {
	if len(tasks) == 0 {
		return tasks, nil
	}

	switch mode {
	case subtasksview.Hidden:
		filtered := funk.Filter(tasks, func(x *models.Task) bool {
			return x.Parent == 0
		})
		tasksNew, _ := filtered.([]*models.Task)
		return tasksNew, nil
	case subtasksview.Indented:
		p2c := make(map[int64][]*models.Task)
		ids, _ := funk.Map(tasks, func(x *models.Task) int64 {
			return x.ID
		}).([]int64)

		topLevels, _ := funk.Filter(tasks, func(x *models.Task) bool {
			// is parent id, or parent is missing
			return x.Parent == 0 || !funk.ContainsInt64(ids, x.Parent)
		}).([]*models.Task)
		children, _ := funk.Filter(tasks, func(x *models.Task) bool {
			return x.Parent != 0
		}).([]*models.Task)
		for _, t := range children {
			p2c[t.Parent] = append(p2c[t.Parent], t)
		}
		tsNew := make([]*models.Task, 0)
		for _, t := range topLevels {
			if iChildren, exist := p2c[t.ID]; exist {
				tsNew = append(tsNew, t)
				tsNew = append(tsNew, iChildren...)
			} else {
				tsNew = append(tsNew, t)
			}
		}

		return tsNew, nil
	case subtasksview.Inline:
		return tasks, nil
	default:
		return nil, fmt.Errorf("unknown mode %d", mode)
	}
}
