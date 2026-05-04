package seleniumtests

import (
	"fmt"
	"log"
	"sync"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	baseURL          = "http://localhost:8080"
	chromeDriverPort = 9515
)

// AppManager — управляет браузером и хелперами (Инструкция 3-4: AppManager + синглтон)
type AppManager struct {
	driver  selenium.WebDriver
	service *selenium.Service
	Auth    *AuthHelper
	Nav     *NavigationHelper
	Todo    *TodoHelper
}

var (
	appInstance *AppManager
	once        sync.Once
)

// GetInstance — возвращает единственный экземпляр AppManager (Инструкция 4: ThreadLocal-аналог)
func GetInstance() *AppManager {
	once.Do(func() {
		service, err := selenium.NewChromeDriverService("chromedriver", chromeDriverPort)
		if err != nil {
			log.Fatalf("Failed to start ChromeDriver: %v", err)
		}

		caps := selenium.Capabilities{"browserName": "chrome"}
		caps.AddChrome(chrome.Capabilities{})

		wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", chromeDriverPort))
		if err != nil {
			service.Stop()
			log.Fatalf("Failed to connect to WebDriver: %v", err)
		}

		wd.MaximizeWindow("")

		appInstance = &AppManager{
			driver:  wd,
			service: service,
		}
		// Инициализируем хелперы через менеджер (Инструкция 3: конструкторы хелперов принимают менеджера)
		appInstance.Auth = &AuthHelper{newHelperBase(appInstance)}
		appInstance.Nav = &NavigationHelper{newHelperBase(appInstance)}
		appInstance.Todo = &TodoHelper{newHelperBase(appInstance)}
	})
	return appInstance
}

// Stop — завершает работу браузера (Инструкция 4: деструктор AppManager)
func (a *AppManager) Stop() {
	if a.driver != nil {
		a.driver.Quit()
	}
	if a.service != nil {
		a.service.Stop()
	}
}
