package tests

import (
	"todo/test6/helpers"
	"todo/test6/models"
)

// AuthBase — базовый класс для тестов, которым нужна предварительная авторизация.
// [SetUp]: вызывает TestBase.Setup() (открывает домашнюю страницу),
// затем выполняет «умную» авторизацию из настроечного файла.
// Тесты, использующие авторизацию как предусловие, наследуются от AuthBase.
type AuthBase struct {
	TestBase
}

func (ab *AuthBase) Setup() {
	ab.TestBase.Setup()
	app.Auth.Login(models.AccountData{
		Username: helpers.Settings.Login,
		Password: helpers.Settings.Password,
	})
}
