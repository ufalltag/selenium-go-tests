// Инструкция 1: запись теста авторизации через рекордер (Katalon/Selenium IDE).
// Всё в одном файле — ровно то, что генерирует рекордер.
// Никаких хелперов, базовых классов и рефакторинга.
package test1

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type untitledTestCase struct {
	driver             selenium.WebDriver
	service            *selenium.Service
	baseURL            string
	verificationErrors []string
	acceptNextAlert    bool
}

func setUp(t *testing.T) *untitledTestCase {
	service, err := selenium.NewChromeDriverService("chromedriver", 9515)
	if err != nil {
		t.Fatalf("start ChromeDriver: %v", err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{})

	wd, err := selenium.NewRemote(caps, "http://localhost:9515/wd/hub")
	if err != nil {
		service.Stop()
		t.Fatalf("connect WebDriver: %v", err)
	}

	wd.SetImplicitWaitTimeout(30 * time.Second)

	return &untitledTestCase{
		driver:             wd,
		service:            service,
		baseURL:            "https://www.google.com/",
		verificationErrors: []string{},
		acceptNextAlert:    true,
	}
}

func (s *untitledTestCase) tearDown(t *testing.T) {
	s.driver.Quit()
	s.service.Stop()
	if len(s.verificationErrors) > 0 {
		t.Errorf("verification errors: %v", s.verificationErrors)
	}
}

func (s *untitledTestCase) isElementPresent(by, value string) bool {
	_, err := s.driver.FindElement(by, value)
	return err == nil
}

// TestUntitledTestCase — прямой перевод test1.py на Go
func TestUntitledTestCase(t *testing.T) {
	s := setUp(t)
	defer s.tearDown(t)

	driver := s.driver
	driver.Get("http://localhost:8080/login")

	username, _ := driver.FindElement(selenium.ByName, "username")
	username.Click()
	username.Clear()
	username.SendKeys("ufalltag")

	password, _ := driver.FindElement(selenium.ByName, "password")
	password.Click()
	password.Clear()
	password.SendKeys("itpavel2005")

	btn, _ := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	btn.Click()

	driver.Get("http://localhost:8080/")
}
