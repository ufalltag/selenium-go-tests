package test5

import (
	"fmt"
	"log"
	"sync"
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

var (
	instance *AppManager
	once     sync.Once
)

func GetInstance() *AppManager {
	once.Do(func() {
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
		instance = &AppManager{driver: wd, service: service}
		instance.Auth = &AuthHelper{HelperBase{manager: instance, driver: wd}}
		instance.Nav = &NavigationHelper{HelperBase{manager: instance, driver: wd}}
		instance.Todo = &TodoHelper{HelperBase{manager: instance, driver: wd}}
	})
	return instance
}

func (a *AppManager) Stop() {
	if a.driver != nil {
		a.driver.Quit()
	}
	if a.service != nil {
		a.service.Stop()
	}
}
