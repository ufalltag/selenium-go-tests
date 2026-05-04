package seleniumtests

// Инструкция 5: тест с загрузкой данных из XML-файла (аналог TestCaseSource в C#/NUnit)
// Сгенерировать данные: go run ./cmd/generator g 3 todos.xml xml
// Затем запустить: go test -run TestCreateTodoFromXML

import (
	"encoding/xml"
	"errors"
	"os"
	"testing"
)

type arrayOfTodoData struct {
	XMLName xml.Name   `xml:"ArrayOfTodoData"`
	Items   []TodoData `xml:"TodoData"`
}

func loadTodosFromXML(filename string) ([]TodoData, error) {
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

func TestCreateTodoFromXML(t *testing.T) {
	todos, err := loadTodosFromXML("todos.xml")
	if errors.Is(err, os.ErrNotExist) {
		t.Skip("todos.xml not found — run: go run ./cmd/generator g 3 todos.xml xml")
	}
	if err != nil {
		t.Fatalf("loadTodosFromXML: %v", err)
	}
	if len(todos) == 0 {
		t.Fatal("todos.xml is empty")
	}

	app.Nav.GoToLoginPage()
	if err := app.Auth.Login(testUser); err != nil {
		t.Fatalf("Login: %v", err)
	}

	// Параметризованный тест — каждый TodoData из XML запускается отдельным подтестом
	for _, todo := range todos {
		todo := todo
		t.Run(todo.Title, func(t *testing.T) {
			if err := app.Todo.CreateTodo(todo); err != nil {
				t.Fatalf("CreateTodo: %v", err)
			}

			// Assert: задача появилась в списке
			present, err := app.Todo.IsTodoPresent(todo.Title)
			if err != nil {
				t.Fatalf("IsTodoPresent: %v", err)
			}
			if !present {
				t.Errorf("Todo %q not found after creation", todo.Title)
			}
		})
	}

	app.Auth.Logout()
}
