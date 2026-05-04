package test2

import "testing"

func TestLogin(t *testing.T) {
	base := &TestBase{}
	base.SetUp(t)
	defer base.TearDown(t)

	user := AccountData{Username: "ufalltag", Password: "itpavel2005"}
	base.GoToLoginPage()
	base.Login(user)
	base.GoToMainPage()
}
