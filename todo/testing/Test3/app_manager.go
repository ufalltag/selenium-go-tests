// Инструкция 3: AppManager управляет драйвером и хелперами.
// На данном этапе НЕ синглтон — каждый тест создаёт свой экземпляр браузера.
package test3

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	baseURL    = "http://localhost:8080"
	driverPort = 9515
)

type AppManager struct {
	driver  selenium.WebDriver
	service *selenium.Service
	Auth    *AuthHelper
	Nav     *NavigationHelper
	Todo    *TodoHelper
}

func NewAppManager() *AppManager {
	service, err := selenium.NewChromeDriverService("chromedriver", driverPort)
	if err != nil {
		log.Fatalf("start ChromeDriver: %v", err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{})

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", driverPort))
	if err != nil {
		service.Stop()
		log.Fatalf("connect WebDriver: %v", err)
	}

	wd.MaximizeWindow("")
	wd.SetImplicitWaitTimeout(10 * time.Second)

	app := &AppManager{driver: wd, service: service}
	// Инструкция 3: хелперы получают менеджера в конструкторе
	app.Auth = &AuthHelper{HelperBase{manager: app, driver: wd}}
	app.Nav = &NavigationHelper{HelperBase{manager: app, driver: wd}}
	app.Todo = &TodoHelper{HelperBase{manager: app, driver: wd}}
	return app
}

func (a *AppManager) Stop() {
	if a.driver != nil {
		a.driver.Quit()
	}
	if a.service != nil {
		a.service.Stop()
	}
}
