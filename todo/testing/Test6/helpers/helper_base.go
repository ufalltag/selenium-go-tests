package helpers

import "github.com/tebeka/selenium"

type HelperBase struct {
	manager *AppManager
	driver  selenium.WebDriver
}
