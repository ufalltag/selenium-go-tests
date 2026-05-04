package test3

import "github.com/tebeka/selenium"

type AuthHelper struct {
	HelperBase
}

func (a *AuthHelper) Register(user AccountData) {
	username, _ := a.driver.FindElement(selenium.ByName, "username")
	username.Clear()
	username.SendKeys(user.Username)

	password, _ := a.driver.FindElement(selenium.ByName, "password")
	password.Clear()
	password.SendKeys(user.Password)

	btn, _ := a.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	btn.Click()
}

func (a *AuthHelper) Login(user AccountData) {
	username, _ := a.driver.FindElement(selenium.ByName, "username")
	username.Clear()
	username.SendKeys(user.Username)

	password, _ := a.driver.FindElement(selenium.ByName, "password")
	password.Clear()
	password.SendKeys(user.Password)

	btn, _ := a.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	btn.Click()
}
