// Инструкция 4: тест создания + тест удаления с Assert'ами
package test4

import "testing"

var testUser = AccountData{Username: "ufalltag", Password: "itpavel2005"}

func TestCreateTodo(t *testing.T) {
	app.Nav.GoToLoginPage()
	app.Auth.Login(testUser)

	todo := TodoData{Title: "Тестовая задача", Deadline: "2026-05-01"}
	app.Todo.CreateTodo(todo)

	// Assert: задача появилась в списке
	if !app.Todo.IsTodoPresent(todo.Title) {
		t.Errorf("Todo %q not found after creation", todo.Title)
	}

	app.Auth.Logout()
}

// TestDeleteTodo — дополнительный тест удаления
func TestDeleteTodo(t *testing.T) {
	app.Nav.GoToLoginPage()
	app.Auth.Login(testUser)

	todo := TodoData{Title: "Задача для удаления", Deadline: "2026-06-01"}
	app.Todo.CreateTodo(todo)

	if !app.Todo.IsTodoPresent(todo.Title) {
		t.Fatal("Todo not found after creation")
	}

	app.Todo.DeleteTodoByTitle(todo.Title)

	// Assert: задача исчезла
	if app.Todo.IsTodoPresent(todo.Title) {
		t.Errorf("Todo %q still present after deletion", todo.Title)
	}

	app.Auth.Logout()
}
