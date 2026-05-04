package seleniumtests

import (
	"testing"
)

var testUser = AccountData{Username: "ufalltag", Password: "itpavel2005"}

func TestCreateTodo(t *testing.T) {
	app.Nav.GoToLoginPage()
	if err := app.Auth.Login(testUser); err != nil {
		t.Fatalf("Login: %v", err)
	}

	todo := TodoData{Title: "Тестовая задача", Deadline: "2026-05-01"}
	if err := app.Todo.CreateTodo(todo); err != nil {
		t.Fatalf("CreateTodo: %v", err)
	}

	// Assert: задача появилась в списке (Инструкция 4)
	present, err := app.Todo.IsTodoPresent(todo.Title)
	if err != nil {
		t.Fatalf("IsTodoPresent: %v", err)
	}
	if !present {
		t.Errorf("Todo %q not found in list after creation", todo.Title)
	}

	app.Auth.Logout()
}

// TestDeleteTodo — тест удаления (Инструкция 4: добавить тест редактирования/удаления)
func TestDeleteTodo(t *testing.T) {
	app.Nav.GoToLoginPage()
	if err := app.Auth.Login(testUser); err != nil {
		t.Fatalf("Login: %v", err)
	}

	// Создаём задачу, которую потом удалим
	todo := TodoData{Title: "Задача для удаления", Deadline: "2026-06-01"}
	if err := app.Todo.CreateTodo(todo); err != nil {
		t.Fatalf("CreateTodo: %v", err)
	}

	present, err := app.Todo.IsTodoPresent(todo.Title)
	if err != nil {
		t.Fatalf("IsTodoPresent: %v", err)
	}
	if !present {
		t.Fatal("Todo not found after creation — cannot proceed with delete test")
	}

	// Удаляем
	if err := app.Todo.DeleteTodoByTitle(todo.Title); err != nil {
		t.Fatalf("DeleteTodoByTitle: %v", err)
	}

	// Assert: задача исчезла из списка
	present, err = app.Todo.IsTodoPresent(todo.Title)
	if err != nil {
		t.Fatalf("IsTodoPresent after delete: %v", err)
	}
	if present {
		t.Errorf("Todo %q still present after deletion", todo.Title)
	}

	app.Auth.Logout()
}
