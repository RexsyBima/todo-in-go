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

type Todo struct {
	Title   string
	Done    bool
	Status  StatusLevel
	Created time.Time
}

func (t Todo) GetDetail() string {
	return fmt.Sprintf("Title : %v | Done : %v | Status : %v | Created %v ", t.Title, t.Done, t.Status, t.Created.Format("02/01/2006 15:04:05 MST"))
}
