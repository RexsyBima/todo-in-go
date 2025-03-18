package utils

import "fmt"

type Menu int

const (
	MenuAdd     Menu = iota
	MenuAddMany Menu = iota
	MenuView    Menu = iota
	MenuUpdate  Menu = iota
	MenuDelete  Menu = iota
	MenuSave    Menu = iota
	MenuQuit    Menu = iota
)

func ShowMenu() {
	println("1. Add a todo")
	println("2. Add many todo")
	println("3. View")
	println("4. Update")
	println("5. Delete")
	println("6. Save")
	println("7. Quit")
}

func ParseUserInput(prompt string) int {
	var input int
	fmt.Println(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("invalid input, please enter a valid integer (from 1 to 7)")
		return -1
	}
	return input
}
