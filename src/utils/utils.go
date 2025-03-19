package utils

import "todo_app/src/model"

func AddTodos(todo model.Todo, slices []model.Todo) []model.Todo {
	return append(slices, todo)
}

func Intro() {

}
