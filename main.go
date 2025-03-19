package main

import (
	// "errors"
	"fmt"
	"time"
	"todo_app/src/model"
	"todo_app/src/utils"
	"todo_app/src/utils/menu"
)

func CreateTodo() model.Todo {
	var todo model.Todo
	fmt.Println("enter todo title: ")
	fmt.Scanln(&todo.Title)
	fmt.Println("enter todo status, 0 for not important and 1 for important: ")
	_, err := fmt.Scanln(&todo.Status)
	if err != nil || todo.Status < 0 || todo.Status > 2 {
		fmt.Println("invalid todo status input, default to 'undefined'")
		todo.Status = 0
	}
	todo.Created = time.Now()
	return todo
}

func main() {
	const firstName string = "Rexsy"
	const lastName string = "Bima"
	const appName string = "Todo App"
	const stackName string = "Golang"
	const versionName string = "go version go1.24.1 linux/amd64"
	var timeNow string = time.Now().Format("02/01/2006 15:04:05 MST")
	menu.ShowIntro(appName, stackName, firstName, lastName, timeNow, versionName)
	menu.ShowMenu()
	var todos []model.Todo
	user_action, _ := menu.ParseUserInput("Please enter a user action: ")
	// TODO: implement the user action, save them to slice
	switch user_action {
	case menu.MenuAdd:
		todo := CreateTodo()
		if todo.Title != "" && todo.Title != "0" {
			todos = utils.AddTodos(todo, todos)
		} else {
			fmt.Println("todo is not added because its title is empty")

		}
	case menu.MenuAddMany:
		for {
			todo := CreateTodo()
			if todo.Title != "" && todo.Title != "0" {
				todos = append(todos, todo)
			}
			var user_input string
			fmt.Println("Press 1 to add more todos, press other to exit")
			fmt.Scanln(&user_input)
			if user_input == "1" {
				continue
			} else {
				break
			}
		}
	case menu.MenuView:
		// TODO: add a load and save data feature function
		for i, v := range todos {
			fmt.Println(i, v.GetDetail())
		}
	case menu.MenuUpdate:
		var user_input int
		if len(todos) == 0 {
			fmt.Println("Todos are empty")
			break
		}
		for i, v := range todos {
			fmt.Println(i, v.GetDetail())
		}
		fmt.Println("Please enter a todo number to update: ")
		_, err := fmt.Scanln(&user_input)
		if err != nil || user_input < 1 || user_input > len(todos) {
			fmt.Printf("Invalid input, please enter a valid integer (from 1 to %v)\n", len(todos))
		}
	case menu.MenuDelete:
		panic("this shit has not been emplemented yet")
	case menu.MenuSave:
		panic("this shit has not been emplemented yet")
	case menu.MenuQuit:
		panic("this shit has not been emplemented yet")
	default:
		panic("this shit has not been emplemented yet")
	}
	// var todo string
	// fmt.Scanln(&todo)
	for i, v := range todos {
		fmt.Println(fmt.Sprintf("%v.", i+1), v.GetDetail())
	}
	todoCount := len(todos)
	fmt.Printf("todo count is %v", todoCount)
	fmt.Println()
}
