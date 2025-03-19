package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"todo_app/src/model"
)

func AddTodos(todo model.Todo, slices []model.Todo) []model.Todo {
	return append(slices, todo)
}

func SaveTodosCSV(todos []model.Todo, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	defer file.Close()
	defer writer.Flush()
	writer.Write([]string{"Title", "Done", "Status", "Created"})
	for _, user := range todos {
		writer.Write([]string{user.Title, strconv.FormatBool(user.Done), strconv.Itoa(int(user.Status)), user.Created.Format("02/01/2006 15:04:05 MST")})
	}
	return nil
}

func LoadTodosCSV(filename string) (t []model.Todo, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
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
		t = append(t, todo)
	}
	return
}
