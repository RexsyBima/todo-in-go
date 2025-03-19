package menu

import "fmt"

type Menu int

const (
	MenuAdd Menu = iota
	MenuAddMany
	MenuMarkDone
	MenuView
	MenuUpdate
	MenuDelete
	MenuDeleteAll
	MenuSave
	MenuQuit
	TotalMenu
)

func ShowMenu() {
	fmt.Printf("%v. Add a todo\n", MenuAdd+1)
	fmt.Printf("%v. Add many todo\n", MenuAddMany+1)
	fmt.Printf("%v. Mark a todo as done (vice versa)\n", MenuMarkDone+1)
	fmt.Printf("%v. View\n", MenuView+1)
	fmt.Printf("%v. Update\n", MenuUpdate+1)
	fmt.Printf("%v. Delete a todo\n", MenuDelete+1)
	fmt.Printf("%v. Delete all todos\n", MenuDeleteAll+1)
	fmt.Printf("%v. Save\n", MenuSave+1)
	fmt.Printf("%v. Quit\n", MenuQuit+1)
}

func ShowIntro(appName, stackName, firstName, lastName, timeNow, version string) {
	fmt.Printf("Welcome to the %v made in %v ...", appName, stackName)
	fmt.Printf("Today's date is %v \n", timeNow)
	fmt.Println("Transitioning from python to go")
	fmt.Printf("created by: %v %v \n", firstName, lastName)
	fmt.Printf("Go version: %v \n", version)
	fmt.Println("Please enter a todo: ")
}

func ParseUserInput(prompt string) (Menu, error) {
	var input int
	fmt.Println(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil || input < int(MenuAdd) || input > int(TotalMenu) {
		fmt.Println("Invalid input, please enter a valid integer (from 1 to 8)")
		return MenuQuit, fmt.Errorf("invalid input: %v", err)
	}
	return Menu(input - 1), nil
}
