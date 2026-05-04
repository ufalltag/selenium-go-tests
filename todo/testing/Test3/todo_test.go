package test3

import "testing"

func TestCreateTodo(t *testing.T) {
	base := &TestBase{}
	base.SetUp()
	defer base.TearDown()

	base.app.Nav.GoToLoginPage()
	base.app.Auth.Login(AccountData{Username: "ufalltag", Password: "itpavel2005"})

	todo := TodoData{Title: "Тестовая задача", Deadline: "2026-05-01"}
	base.app.Todo.CreateTodo(todo)
}
