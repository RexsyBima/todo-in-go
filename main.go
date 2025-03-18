package main

import (
	"fmt"
	"time"
	"todo_app/src/utils"
)

func main() {
	const firstName string = "Rexsy"
	const lastName string = "Bima"
	const appName string = "Todo App"
	const stackName string = "Golang"
	const versionName string = "go version go1.24.1 linux/amd64"
	var timeNow string = time.Now().String()
	const todoCount uint8 = 50
	fmt.Printf("Welcome to the %v made in %v ...", appName, stackName)
	fmt.Printf("Today's date is %v \n", timeNow)
	fmt.Println("Transitioning from python to go")
	fmt.Printf("created by: %v %v \n", firstName, lastName)
	fmt.Println("Please enter a todo: ")
	var todos []string
	utils.ShowMenu()
	user_action := utils.ParseUserInput("Please enter a user action: ")
	// TODO: implement the user action, save them to slice
	switch user_action {
	case 1:
		panic("this shit has not been emplemented yet")
	case 2:
		panic("this shit has not been emplemented yet")
	case 3:
		panic("this shit has not been emplemented yet")
	case 4:
		panic("this shit has not been emplemented yet")
	default:
		panic("this shit has not been emplemented yet")
	}
	var todo string
	fmt.Scanln(&todo)
	if todo != "" && todo != "0" {
		todos = append(todos, todo)
	}
	for i := range todos {
		fmt.Println(todos[i])
	}
	var remainingTodo uint8 = 10
	fmt.Printf("todo count is %v and its remainingTodo is %v \n", todoCount, remainingTodo)
}
