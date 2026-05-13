package tests

import (
	"encoding/xml"
	"errors"
	"os"
	"testing"
	"todo/test6/models"
)

type arrayOfTodoData struct {
	XMLName xml.Name          `xml:"ArrayOfTodoData"`
	Items   []models.TodoData `xml:"TodoData"`
}

func loadTodosFromXML(filename string) ([]models.TodoData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var arr arrayOfTodoData
	if err := xml.Unmarshal(data, &arr); err != nil {
		return nil, err
	}
	return arr.Items, nil
}

var xmlBase AuthBase

func TestCreateTodoFromXML(t *testing.T) {
	todos, err := loadTodosFromXML("todos.xml")
	if errors.Is(err, os.ErrNotExist) {
		t.Skip("todos.xml not found — run: go run ../cmd/generator g 3 todos.xml xml")
	}
	if err != nil {
		t.Fatalf("loadTodosFromXML: %v", err)
	}

	xmlBase.Setup()

	for _, todo := range todos {
		todo := todo
		t.Run(todo.Title, func(t *testing.T) {
			app.Todo.CreateTodo(todo)

			found, ok := app.Todo.FindTodo(todo.Title)
			if !ok {
				t.Fatalf("todo %q not found after creation", todo.Title)
			}

			assertEqualStr(t, "Title", todo.Title, found.Title)
			assertEqualStr(t, "Deadline", toDisplayDate(todo.Deadline), found.Deadline)
			if found.ID == "" {
				t.Errorf("ID: expected non-empty DB id")
			}
		})
	}
}
