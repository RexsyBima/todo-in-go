package logic

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type Menu int

const (
	MenuAdd    Menu = iota
	MenuView   Menu = iota
	MenuUpdate Menu = iota
	MenuDelete Menu = iota
)

func TestMenu(t *testing.T) {
	switch Menu(1 - 1) {
	case MenuAdd:
		fmt.Println("Add")
	case MenuView:
		fmt.Println("View")
	case MenuUpdate:
		fmt.Println("Update")
	case MenuDelete:
		fmt.Println("Delete")
	}
}

func TestSlices(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	var arr2 = []int{6, 7, 8, 9, 10}
	arr3 := append(arr, arr2...)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(arr3, expected) {
		t.Errorf("expected %v but got %v", expected, arr3)
	}
}

func TestEquivalent(t *testing.T) {
	x := ""
	var y string
	fmt.Printf("%v", x == y)
}

func TestWriteData(t *testing.T) {
	name := "rexxxx"
	filename := "output.txt"
	file, _ := os.Create(filename)
	fmt.Fprint(file, name)
}
