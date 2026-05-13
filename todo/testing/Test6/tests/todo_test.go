package tests

import (
	"testing"
	"todo/test6/models"
)

var todoBase AuthBase

func TestCreateTodo(t *testing.T) {
	todoBase.Setup()

	// UUID в заголовке — каждый запуск создаёт уникальную задачу,
	// исключая ложные срабатывания от данных предыдущих запусков.
	todo := models.TodoData{
		Title:    "Тестовая задача " + newID(),
		Deadline: "2026-05-01",
	}
	app.Todo.CreateTodo(todo)

	found, ok := app.Todo.FindTodo(todo.Title)
	if !ok {
		t.Fatalf("todo %q not found after creation", todo.Title)
	}

	// assertEqual на каждое поле — не просто факт наличия
	assertEqualStr(t, "Title", todo.Title, found.Title)
	assertEqualStr(t, "Deadline", toDisplayDate(todo.Deadline), found.Deadline)
	if found.ID == "" {
		t.Errorf("ID: expected non-empty DB id")
	}
}

func TestDeleteTodo(t *testing.T) {
	todoBase.Setup()

	todo := models.TodoData{
		Title:    "Задача для удаления " + newID(),
		Deadline: "2026-06-01",
	}
	app.Todo.CreateTodo(todo)

	found, ok := app.Todo.FindTodo(todo.Title)
	if !ok {
		t.Fatal("todo not found after creation")
	}
	assertEqualStr(t, "Title", todo.Title, found.Title)
	assertEqualStr(t, "Deadline", toDisplayDate(todo.Deadline), found.Deadline)

	// Удаляем по ID из БД — точнее чем поиск по заголовку
	app.Todo.DeleteTodo(found.ID)

	if app.Todo.IsTodoPresent(todo.Title) {
		t.Errorf("todo %q still present after deletion", todo.Title)
	}
}
