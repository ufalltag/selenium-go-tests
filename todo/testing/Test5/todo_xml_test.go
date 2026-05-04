// Инструкция 5: параметризованный тест — данные загружаются из XML-файла.
// Аналог [TestCaseSource] в C#/NUnit.
//
// Как использовать:
//  1. go run ./cmd/generator g 3 todos.xml xml
//  2. go test -v -run TestCreateTodoFromXML
package test5

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

// loadTodosFromXML — аналог метода GroupDataFromXmlFile() из C# примера в ТЗ
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

	app.Nav.GoToLoginPage()
	app.Auth.Login(AccountData{Username: "ufalltag", Password: "itpavel2005"})

	// Каждый объект из XML — отдельный подтест (аналог TestCaseSource)
	for _, todo := range todos {
		todo := todo
		t.Run(todo.Title, func(t *testing.T) {
			app.Todo.CreateTodo(todo)

			if !app.Todo.IsTodoPresent(todo.Title) {
				t.Errorf("Todo %q not found after creation", todo.Title)
			}
		})
	}

	app.Auth.Logout()
}
