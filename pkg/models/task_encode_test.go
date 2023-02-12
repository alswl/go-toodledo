package models

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
)

type Dog struct {
	Name *string `json:"name,omitempty"`
}

func TestPrettyYAML(t *testing.T) {
	task := Task{
		Added:       1642507200,
		Addedby:     123,
		Attatchment: "",
		Children:    0,
		Completed:   0,
		Context:     1292483,
		Duedate:     1646308800,
		Duedatemod:  0,
		Duetime:     0,
		Folder:      9169511,
		Goal:        511079,
		ID:          327077755,
		Length:      60,
		Location:    0,
		Meta:        "",
		Modified:    1676044739,
		Note:        "",
		Order:       0,
		Parent:      0,
		Previous:    0,
		Priority:    2,
		Ref:         "",
		Remind:      (0),
		Repeat:      "FREQ=WEEKLY",
		Shared:      0,
		Sharedowner: 0,
		Sharedwith:  0,
		Star:        0,
		Startdate:   1645531200,
		Starttime:   0,
		Status:      1,
		Tag:         "",
		Timer:       0,
		Timeron:     0,
		Title:       "next-action item",
		Via:         "",
	}
	bs := PrettyYAML(&task)
	assert.Equal(t, `added: "2022-01-18 20:00:00"
addedby: 123
children: 0
completed: ""
context: 1292483
duedate: "2022-03-03"
duedatemod: 0
duetime: ""
folder: 9169511
goal: 511079
id: 327077755
length: 60
location: 0
meta: ""
modified: "2023-02-10 23:58:59"
note: ""
order: 0
parent: 0
previous: 0
priority: 2
ref: ""
remind: 0
repeat: FREQ=WEEKLY
sharedowner: 0
sharedwith: 0
startdate: "2022-02-22"
starttime: ""
status: 1
tag: ""
timer: 0
timeron: ""
title: next-action item
`, bs)
}

func TestYAMLForPointerString(t *testing.T) {
	d := Dog{}
	bs, _ := yaml.Marshal(d)
	assert.Equal(t, "name: null\n", string(bs))
}

func TestLoadTaskFromYAML(t *testing.T) {
	var yaml = `added: "2022-01-18 20:00:00"
addedby: 123
children: 0
completed: ""
context: 1292483
duedate: "2022-03-03"
duedatemod: 0
duetime: ""
folder: 9169511
goal: 511079
id: 327077755
length: 60
location: 0
meta: ""
modified: "2023-02-10 23:58:59"
note: ""
order: 0
parent: 0
previous: 0
priority: 2
ref: ""
remind: 0
repeat: FREQ=WEEKLY
sharedowner: 0
sharedwith: 0
startdate: "2022-02-22"
starttime: ""
status: 1
tag: ""
timer: 0
timeron: ""
title: next-action item
`
	task, err := LoadTaskFromYAML(yaml)
	assert.NoError(t, err)

	assert.Equal(t, int64(1642536000), *task.Added)
	assert.Equal(t, int64(123), *task.Addedby)
	assert.Equal(t, int64(0), *task.Children)
	assert.Equal(t, int64(0), *task.Completed)
	assert.Equal(t, int64(1292483), *task.Context)
	assert.Equal(t, int64(1646265600), *task.Duedate)
	assert.Equal(t, int64(0), *task.Duedatemod)
	assert.Equal(t, int64(0), *task.Duetime)
	assert.Equal(t, int64(9169511), task.Folder)
}
