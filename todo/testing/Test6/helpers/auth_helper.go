package helpers

import (
	"todo/test6/models"

	"github.com/tebeka/selenium"
)

type AuthHelper struct {
	HelperBase
	currentUser string
}

// IsLoggedIn без аргумента — залогинен ли кто-либо.
// IsLoggedIn(username) — залогинен ли именно этот пользователь.
// Аналог двух перегрузок IsLoggedIn() и IsLoggedIn(string) из C#.
func (a *AuthHelper) IsLoggedIn(username ...string) bool {
	_, err := a.driver.FindElement(selenium.ByLinkText, "Выйти")
	loggedIn := err == nil
	if len(username) == 0 {
		return loggedIn
	}
	return loggedIn && a.currentUser == username[0]
}

// Login — «умная» авторизация: не логинится повторно если уже под нужным пользователем,
// и не логаутится если уже разлогинены.
func (a *AuthHelper) Login(user models.AccountData) {
	if a.IsLoggedIn() {
		if a.IsLoggedIn(user.Username) {
			return
		}
		a.Logout()
	}
	a.driver.Get(Settings.BaseURL + "/login")

	username, _ := a.driver.FindElement(selenium.ByName, "username")
	username.Clear()
	username.SendKeys(user.Username)

	password, _ := a.driver.FindElement(selenium.ByName, "password")
	password.Clear()
	password.SendKeys(user.Password)

	btn, _ := a.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	btn.Click()

	a.currentUser = user.Username
}

// Logout — «умный» выход: не пытается выйти если уже разлогинены.
func (a *AuthHelper) Logout() {
	if !a.IsLoggedIn() {
		return
	}
	link, _ := a.driver.FindElement(selenium.ByLinkText, "Выйти")
	link.Click()
	a.currentUser = ""
}
