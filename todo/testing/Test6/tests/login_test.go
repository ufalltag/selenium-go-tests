package tests

import (
	"testing"
	"todo/test6/helpers"
	"todo/test6/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var loginBase TestBase

func TestLoginWithValidData(t *testing.T) {
	loginBase.Setup()
	app.Auth.Logout()

	app.Nav.GoToLoginPage()
	app.Auth.Login(models.AccountData{
		Username: helpers.Settings.Login,
		Password: helpers.Settings.Password,
	})

	assert.True(t, app.Nav.IsOnMainPage(), "Expected main page after valid login")
}

func TestLoginWithInvalidData(t *testing.T) {
	loginBase.Setup()
	app.Auth.Logout()

	app.Nav.GoToLoginPage()
	app.Auth.Login(models.AccountData{Username: "wrong_user", Password: "wrong_pass"})

	url, err := app.Nav.CurrentURL()
	require.NoError(t, err)
	assert.Contains(t, url, "/login", "Expected to stay on login page")
}
