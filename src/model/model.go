package model

import (
	"fmt"
	"time"
)

type StatusLevel int

const (
	Undefined StatusLevel = iota
	NotImportant
	Important
)

type Owner struct {
	Name  string
	Admin bool
}

type Todo struct {
	Title   string
	Done    bool
	Status  StatusLevel
	Created time.Time
	Owner
}

func (t Todo) GetDetail() string {
	return fmt.Sprintf("User : %v | Admin : %v | Title : %v | Done : %v | Status : %v | Created %v ", t.Name, t.Admin, t.Title, t.Done, t.Status, t.Created.Format("02/01/2006 15:04:05 MST"))
}

func (t *Todo) Update(title string, done bool, status StatusLevel) {
	t.Title = title
	t.Done = done
	t.Status = status

}
