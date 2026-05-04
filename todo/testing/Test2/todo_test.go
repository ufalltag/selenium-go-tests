// Инструкция 2: второй тест — добавление сущности (задачи)
package test2

import "testing"

func TestCreateTodo(t *testing.T) {
	base := &TestBase{}
	base.SetUp(t)
	defer base.TearDown(t)

	base.GoToLoginPage()
	base.Login(AccountData{Username: "ufalltag", Password: "itpavel2005"})

	todo := TodoData{Title: "Тестовая", Deadline: "2026-05-01"}
	base.CreateTodo(todo)
}
