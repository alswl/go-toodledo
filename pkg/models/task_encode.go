package models

import (
	"time"

	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/jinzhu/copier"
	"sigs.k8s.io/yaml"
)

// TaskEditableFormatted is the task that can be edited.
// all the date time fields will convert to human-readable format.
type TaskEditableFormatted struct {
	TaskEdit

	// Added is the time when the task is added.
	// it was NO omitempty, we should display in editor so that user can edit it.
	// it was skip copier, it should be convert format
	// and other fields are same as Added
	Added     string `json:"added" copier:"-"`
	Completed string `json:"completed" copier:"-"`
	Duedate   string `json:"duedate" copier:"-"`
	Duetime   string `json:"duetime" copier:"-"`
	Modified  string `json:"modified" copier:"-"`
	Startdate string `json:"startdate" copier:"-"`
	Starttime string `json:"starttime" copier:"-"`
	Timeron   string `json:"timeron" copier:"-"`
}

func PrettyYAML(t *Task) string {
	var newT = TaskEditableFormatted{}
	_ = copier.Copy(&newT, t)

	// TODO check is added is date only or date and time
	if t.Added != 0 {
		added := time.Unix(t.Added, 0)
		newT.Added = added.In(utils.DefaultTimeZone).Format(constants.DefaultTimeLayout)
	}
	if t.Completed != 0 {
		completed := time.Unix(t.Completed, 0)
		newT.Completed = (completed.In(utils.DefaultTimeZone).Format(constants.DefaultTimeLayout))
	}
	if t.Duedate != 0 {
		duedate := time.Unix(t.Duedate, 0)
		// duedate is calculated in UTC(ignore timezone)
		newT.Duedate = (duedate.Format(constants.DefaultDateOnlyLayout))
	}
	if t.Duetime != 0 {
		duetime := time.Unix(t.Duetime, 0)
		// duetime is calculated in UTC(ignore timezone)
		newT.Duetime = (duetime.Format(constants.DefaultTimeOnlyLayout))
	}
	if t.Modified != 0 {
		modified := time.Unix(t.Modified, 0)
		newT.Modified = (modified.In(utils.DefaultTimeZone).Format(constants.DefaultTimeLayout))
	}
	if t.Startdate != 0 {
		startdate := time.Unix(t.Startdate, 0)
		// startdate is calculated in UTC(ignore timezone)
		newT.Startdate = (startdate.Format(constants.DefaultDateOnlyLayout))
	}
	if t.Starttime != 0 {
		starttime := time.Unix(t.Starttime, 0)
		newT.Starttime = (starttime.In(utils.DefaultTimeZone).Format(constants.DefaultTimeLayout))
	}
	if t.Timeron != 0 {
		timeron := time.Unix(t.Timeron, 0)
		newT.Timeron = (timeron.Format(constants.DefaultTimeOnlyLayout))
	}

	bs, _ := yaml.Marshal(newT)
	return string(bs)
}

func LoadTaskFromYAML(prettyFormatYAMLString string) (*TaskEdit, error) {
	t := &TaskEditableFormatted{}
	newT := &TaskEdit{}
	err := yaml.Unmarshal([]byte(prettyFormatYAMLString), t)
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(newT, t)

	if t.Added != "" {
		added, ierr := time.Parse(constants.DefaultTimeLayout, t.Added)
		if ierr != nil {
			return nil, ierr
		}
		newT.Added = utils.WrapPointerInt64(added.Unix())
	}
	if t.Completed != "" {
		completed, ierr := time.Parse(constants.DefaultTimeLayout, t.Completed)
		if ierr != nil {
			return nil, ierr
		}
		newT.Completed = utils.WrapPointerInt64(completed.Unix())
	}
	if t.Duedate != "" {
		duedate, ierr := time.Parse(constants.DefaultDateOnlyLayout, t.Duedate)
		if ierr != nil {
			return nil, ierr
		}
		newT.Duedate = utils.WrapPointerInt64(duedate.Unix())
	}
	if t.Duetime != "" {
		// fill 1970-01-01 to make it a valid time
		duetime, ierr := time.Parse(constants.DefaultTimeLayout, "1970-01-01 "+t.Duetime+":00")
		if ierr != nil {
			return nil, ierr
		}
		newT.Duetime = utils.WrapPointerInt64(duetime.Unix())
	}
	if t.Modified != "" {
		modified, ierr := time.Parse(constants.DefaultTimeLayout, t.Modified)
		if ierr != nil {
			return nil, ierr
		}
		newT.Modified = utils.WrapPointerInt64(modified.Unix())
	}
	if t.Startdate != "" {
		startdate, ierr := time.Parse(constants.DefaultDateOnlyLayout, t.Startdate)
		if ierr != nil {
			return nil, ierr
		}
		newT.Startdate = utils.WrapPointerInt64(startdate.Unix())
	}
	if t.Starttime != "" {
		starttime, ierr := time.Parse(constants.DefaultTimeLayout, t.Starttime)
		if ierr != nil {
			return nil, ierr
		}
		newT.Starttime = utils.WrapPointerInt64(starttime.Unix())
	}
	if t.Timeron != "" {
		timeron, ierr := time.Parse(constants.DefaultTimeOnlyLayout, t.Timeron)
		if ierr != nil {
			return nil, ierr
		}
		newT.Timeron = utils.WrapPointerInt64(timeron.Unix())
	}

	return newT, nil
}
