package tests

import (
	"os"
	"testing"
	"todo/test6/helpers"
)

var app *helpers.AppManager

func TestMain(m *testing.M) {
	app = helpers.GetInstance()
	code := m.Run()
	app.Stop()
	os.Exit(code)
}
