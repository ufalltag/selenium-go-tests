package seleniumtests

import "github.com/tebeka/selenium"

type AuthHelper struct {
	HelperBase
}

func (a *AuthHelper) Login(user AccountData) error {
	usernameField, err := a.driver.FindElement(selenium.ByName, "username")
	if err != nil {
		return err
	}
	usernameField.Clear()
	usernameField.SendKeys(user.Username)

	passwordField, err := a.driver.FindElement(selenium.ByName, "password")
	if err != nil {
		return err
	}
	passwordField.Clear()
	passwordField.SendKeys(user.Password)

	submitBtn, err := a.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	if err != nil {
		return err
	}
	return submitBtn.Click()
}

func (a *AuthHelper) Register(user AccountData) error {
	usernameField, err := a.driver.FindElement(selenium.ByName, "username")
	if err != nil {
		return err
	}
	usernameField.Clear()
	usernameField.SendKeys(user.Username)

	passwordField, err := a.driver.FindElement(selenium.ByName, "password")
	if err != nil {
		return err
	}
	passwordField.Clear()
	passwordField.SendKeys(user.Password)

	submitBtn, err := a.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	if err != nil {
		return err
	}
	return submitBtn.Click()
}

func (a *AuthHelper) Logout() error {
	logoutLink, err := a.driver.FindElement(selenium.ByLinkText, "Выйти")
	if err != nil {
		return err
	}
	return logoutLink.Click()
}
