package models

import "fmt"

func (t *Task) String() string {
	return fmt.Sprintf("%d %s", t.ID, t.Title)
}
