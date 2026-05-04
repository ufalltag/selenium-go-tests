package test4

import (
	"os"
	"testing"
)

var app *AppManager

func TestMain(m *testing.M) {
	app = GetInstance()
	code := m.Run()
	app.Stop()
	os.Exit(code)
}
