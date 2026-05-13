package tests

import "todo/test6/helpers"

// TestBase — базовый класс для всех тестов.
// [SetUp]: получает instance менеджера и открывает домашнюю страницу.
// Тесты, проверяющие авторизацию, наследуются от TestBase напрямую.
type TestBase struct{}

func (tb *TestBase) Setup() {
	app = helpers.GetInstance()
	app.Nav.OpenHomePage()
}
