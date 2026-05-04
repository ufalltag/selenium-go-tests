// Инструкция 4: тесты с Assert'ами после каждого действия
package test4

import (
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	app.Nav.GoToLoginPage()
	app.Auth.Login(AccountData{Username: "ufalltag", Password: "itpavel2005"})

	// Assert: ждём элемент главной страницы через implicit wait (10s).
	// Надёжнее чем сразу проверять URL — редирект может ещё не завершиться.
	if !app.Nav.IsOnMainPage() {
		t.Errorf("Expected main page after login")
	}

	app.Auth.Logout()
}

func TestLoginInvalidCredentials(t *testing.T) {
	app.Nav.GoToLoginPage()
	app.Auth.Login(AccountData{Username: "wrong", Password: "wrong"})

	// Assert: остались на странице логина
	url, _ := app.Nav.CurrentURL()
	if !strings.Contains(url, "/login") {
		t.Errorf("Expected login page, got: %s", url)
	}
}
