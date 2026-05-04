package seleniumtests

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestLogin — перевод test1.py на Go (Инструкция 1)
func TestLogin(t *testing.T) {
	if err := app.Nav.GoToLoginPage(); err != nil {
		t.Fatalf("GoToLoginPage: %v", err)
	}

	err := app.Auth.Login(AccountData{
		Username: "ufalltag",
		Password: "itpavel2005",
	})
	if err != nil {
		t.Fatalf("Login: %v", err)
	}

	// Assert: после входа должны быть на главной странице (Инструкция 4)
	url, err := app.Nav.CurrentURL()
	if err != nil {
		t.Fatalf("CurrentURL: %v", err)
	}
	if !strings.HasSuffix(url, "/") {
		t.Errorf("Expected main page after login, got: %s", url)
	}

	app.Auth.Logout()
}

// TestLoginInvalidCredentials — проверка неверных данных (Инструкция 4: Assert)
func TestLoginInvalidCredentials(t *testing.T) {
	if err := app.Nav.GoToLoginPage(); err != nil {
		t.Fatalf("GoToLoginPage: %v", err)
	}

	app.Auth.Login(AccountData{
		Username: "wronguser",
		Password: "wrongpass",
	})

	// Assert: должны остаться на странице логина
	url, err := app.Nav.CurrentURL()
	if err != nil {
		t.Fatalf("CurrentURL: %v", err)
	}
	if !strings.Contains(url, "/login") {
		t.Errorf("Expected to stay on login page, got: %s", url)
	}
}

// TestRegister — уникальный логин через timestamp, чтобы тест работал повторно (Инструкция 2)
func TestRegister(t *testing.T) {
	if err := app.Nav.GoToRegisterPage(); err != nil {
		t.Fatalf("GoToRegisterPage: %v", err)
	}

	username := fmt.Sprintf("user_%d", time.Now().UnixMilli())
	err := app.Auth.Register(AccountData{
		Username: username,
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	// Assert: после регистрации должны быть на главной странице
	url, err := app.Nav.CurrentURL()
	if err != nil {
		t.Fatalf("CurrentURL: %v", err)
	}
	if !strings.HasSuffix(url, "/") {
		t.Errorf("Expected main page after register, got: %s", url)
	}

	app.Auth.Logout()
}
