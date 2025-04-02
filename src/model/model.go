package model

import (
	"fmt"
	"slices"
	"time"
)

type StatusLevel int
type Info interface {
	GetDetail() string
}

func GetDetail(i Info) string {
	return i.GetDetail()
}

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

type Todos []Todo

func (t Todos) ValidateIndex(index int) bool {
	return index >= 0 && index < len(t)
}

func (t *Todos) DeleteTodo(index int) {
	if !t.ValidateIndex(index) {
		return
	}
	*t = slices.Delete(*t, index, index+1)
}

func (t Todo) GetDetail() string {
	return fmt.Sprintf("User : %v | Admin : %v | Title : %v | Done : %v | Status : %v | Created %v ", t.Name, t.Admin, t.Title, t.Done, t.Status, t.Created.Format("02/01/2006 15:04:05 MST"))
}

func (t *Todo) Update(title string, done bool, status StatusLevel) {
	t.Title = title
	t.Done = done
	t.Status = status

}
