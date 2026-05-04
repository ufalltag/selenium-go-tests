package seleniumtests

import (
	"os"
	"testing"
)

var app *AppManager

func TestMain(m *testing.M) {
	app = GetInstance()
	// Инструкция 4: один браузер на все тесты — создаём тестового пользователя перед запуском
	ensureTestUser(AccountData{Username: "ufalltag", Password: "itpavel2005"})
	code := m.Run()
	app.Stop()
	os.Exit(code)
}

// ensureTestUser регистрирует пользователя, если его ещё нет (игнорирует ошибку "уже занято")
func ensureTestUser(user AccountData) {
	app.Nav.GoToRegisterPage()
	app.Auth.Register(user)
	// Если регистрация прошла — выходим. Если нет (уже существует) — тоже нормально.
	app.Auth.Logout()
}
