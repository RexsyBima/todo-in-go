package main

import (
	// "errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
	"todo_app/src/model"
	"todo_app/src/utils"
	"todo_app/src/utils/menu"
)

func CreateTodo(title string, done bool, status model.StatusLevel, owner model.Owner) model.Todo {
	return model.Todo{Title: title, Done: done, Status: status, Created: time.Now(), Owner: owner}
}

func CreateName(i string) string {
	return i
}

func main() {
	var versionNamefunc func(string) string
	const firstName string = "Rexsy"
	const lastName string = "Bima"
	const appName string = "Todo App"
	const stackName string = "Golang"
	user := model.Owner{Name: firstName, Admin: true}
	versionNamefunc = CreateName
	versionName := versionNamefunc("1.0.0")
	var timeNow string = time.Now().Format("02/01/2006 15:04:05 MST")
	menu.ShowIntro(appName, stackName, firstName, lastName, timeNow, versionName)
	todos, _ := utils.LoadTodosCSV("todo_app.csv")
	for {
		menu.ShowMenu()
		user_action, _ := menu.ParseUserInput("Please enter a user action: ")
		// TODO: implement the user action, save them to slice
		switch user_action {
		case menu.MenuAdd:
			var title string
			var status int
			fmt.Println("enter todo title: ")
			fmt.Scanln(&title)
			fmt.Println("enter todo status, 0 for not important and 1 for important: ")
			_, err := fmt.Scanln(&status)
			if err != nil || status < 0 || status > 2 {
				fmt.Println("invalid todo status input, default to 'undefined'")
				status = 0
			}
			todo := CreateTodo(title, false, model.StatusLevel(status), user)
			if todo.Title != "" && todo.Title != "0" {
				todos = utils.AddTodos(todo, todos)
			} else {
				fmt.Println("todo is not added because its title is empty")
			}
		case menu.MenuAddMany:
			for {
				var title string
				var status int
				fmt.Println("enter todo title: ")
				fmt.Scanln(&title)
				fmt.Println("enter todo status, 0 for not important and 1 for important: ")
				_, err := fmt.Scanln(&status)
				if err != nil || status < 0 || status > 2 {
					fmt.Println("invalid todo status input, default to 'undefined'")
					status = 0
				}
				todo := CreateTodo(title, false, model.StatusLevel(status), user)
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
		case menu.MenuMarkDone:
			var user_input int
			if len(todos) == 0 {
				fmt.Println("Todos are empty")
				break
			}
			for i, v := range todos {
				fmt.Println(fmt.Sprintf("%v.", i+1), v.GetDetail())
			}
			fmt.Println("Please enter a todo number to mark it done (vice versa): ")
			_, err := fmt.Scanln(&user_input)
			if err != nil || user_input < 1 || user_input > len(todos) {
				fmt.Printf("Invalid input, please enter a valid integer (from 1 to %v)\n", len(todos))
				break
			}
			switch todos[user_input-1].Done {
			case true:
				todos[user_input-1].Done = false
			case false:
				todos[user_input-1].Done = true
			}
			fmt.Printf("Todo %v has been updated\n", user_input)
		case menu.MenuView:
			fmt.Println("-----------------------------------------------")
			for i, v := range todos {
				func(i int, v model.Todo) {
					fmt.Println(fmt.Sprintf("%v.", i+1), v.GetDetail())
				}(i, v)
			}
			fmt.Println("-----------------------------------------------")
		case menu.MenuUpdate:
			var user_input int
			if len(todos) == 0 {
				fmt.Println("Todos are empty")
				break
			}
			for i, v := range todos {
				fmt.Println(fmt.Sprintf("%v.", i+1), v.GetDetail())
			}
			fmt.Println("Please enter a todo number to update: ")
			_, err := fmt.Scanln(&user_input)
			if err != nil || user_input < 1 || user_input > len(todos) {
				fmt.Printf("Invalid input, please enter a valid integer (from 1 to %v)\n", len(todos))
				break
			}

			var title string
			var status int
			var done string
			fmt.Println("enter todo title: ")
			fmt.Scanln(&title)
			fmt.Println("enter is done?")
			fmt.Scanln(&done)
			fmt.Println("enter todo status, 0 for not important and 1 for important: ")
			_, err = fmt.Scanln(&status)
			if err != nil || status < 0 || status > 2 {
				fmt.Println("invalid todo status input, default to 'undefined'")
				status = 0
			}
			// todo := CreateTodo(title, false, model.StatusLevel(status), user)
			// todo := todos[user_input-1]
			doneConvert, _ := strconv.ParseBool(done)
			todos[user_input-1].Update(title, doneConvert, model.Important)
			fmt.Printf("Todo %v has been updated\n", user_input)
			todoCount := len(todos)
			fmt.Printf("todo count is %v", todoCount)
		case menu.MenuDelete:
			var user_input int
			for i, v := range todos {
				fmt.Println(fmt.Sprintf("%v.", i+1), v.GetDetail())
			}
			fmt.Printf("Please enter a todo number to delete: ")
			_, err := fmt.Scanln(&user_input)
			if err != nil || user_input < 1 || user_input > len(todos) {
				fmt.Printf("Invalid input, please enter a valid integer (from 1 to %v)\n", len(todos))
				break
			}
			todos = slices.Delete(todos, user_input-1, user_input)
			fmt.Printf("Todo %v has been deleted\n", user_input)
			todoCount := len(todos)
			fmt.Printf("todo count is %v", todoCount)
		case menu.MenuDeleteAll:
			if len(todos) == 0 {
				fmt.Printf("Todos are empty")
				break
			}
			todos = make([]model.Todo, 0)
			fmt.Println("all todos has been cleared")
			fmt.Printf("todo count is %v", len(todos))
			fmt.Printf("%v", todos)
		case menu.MenuSave:
			if len(todos) < 1 {
				fmt.Println("Todos are empty")
				break
			}
			utils.SaveTodosCSV(todos, "todo_app.csv")
			fmt.Println("all todos has been saved to todo_app.csv")
		case menu.MenuQuit:
			os.Exit(1)
		default:
			fmt.Println("invalid input")
		}
	}
}
