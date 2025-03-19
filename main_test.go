package main

import (
	// "fmt"
	"slices"
	"testing"
	"time"
	"todo_app/src/model"
)

func TestUpdateTodo(t *testing.T) {
	expected := "updated"
	var todos []model.Todo
	todos = append(todos, model.Todo{"test", false, model.NotImportant, time.Now()})
	var user_input int
	user_input = 1
	name := "updated"
	todos[user_input-1].Title = name
	for _, v := range todos {
		if v.Title != expected {
			t.Errorf("expected %v but got %v", name, v.Title)
		}
	}
}

func TestDeleteATodo(t *testing.T) {
	var todos []model.Todo
	todos = append(todos, model.Todo{"test1", false, model.NotImportant, time.Now()})
	todos = append(todos, model.Todo{"test2", false, model.NotImportant, time.Now()})
	expected := 1
	var user_input int
	user_input = 1
	todos = slices.Delete(todos, user_input, user_input+1)
	if len(todos) != expected {
		t.Errorf("expected %v but got %v", expected, len(todos))
	}

}
