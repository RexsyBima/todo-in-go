package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
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

func TestSaveTodosCSV(t *testing.T) {
	var todos []model.Todo
	todos = append(todos, model.Todo{"test1", false, model.NotImportant, time.Now()})
	todos = append(todos, model.Todo{"test2", false, model.NotImportant, time.Now()})
	file, err := os.Create("output.csv")
	if err != nil {
		t.Errorf("failed to create file: %v", err)
	}
	writer := csv.NewWriter(file)
	defer file.Close()
	defer writer.Flush()
	writer.Write([]string{"Title", "Done", "Status", "Created"})
	for _, user := range todos {
		writer.Write([]string{user.Title, strconv.FormatBool(user.Done), strconv.Itoa(int(user.Status)), user.Created.Format("02/01/2006 15:04:05 MST")})
	}

}

func TestLoadTodosCSV(t *testing.T) {
	file, err := os.Open("output.csv")
	if err != nil {
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}
	var todos []model.Todo
	for i, row := range records {
		if i == 0 {
			continue
		}
		isDone, _ := strconv.ParseBool(row[1])
		statusInt, _ := strconv.Atoi(row[2])
		status := model.StatusLevel(statusInt)
		dateString := row[3]
		convertedTime, _ := time.Parse("02/01/2006 15:04:05 MST", dateString)
		todo := model.Todo{
			Title:   row[0],
			Done:    isDone,
			Status:  status,
			Created: convertedTime,
		}
		todos = append(todos, todo)
	}
	for _, todo := range todos {
		fmt.Println(todo.GetDetail())
	}
}
