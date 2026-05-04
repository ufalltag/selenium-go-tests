package test3

import "testing"

func TestLogin(t *testing.T) {
	base := &TestBase{}
	base.SetUp()
	defer base.TearDown()

	base.app.Nav.GoToLoginPage()
	base.app.Auth.Login(AccountData{Username: "ufalltag", Password: "itpavel2005"})
	base.app.Nav.GoToMainPage()
}
