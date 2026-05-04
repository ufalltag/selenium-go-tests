// Инструкция 2: базовый класс для всех тестов.
// Содержит setUp/tearDown и вспомогательные методы, вынесенные из теста.
package test2

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const baseURL = "http://localhost:8080"

type TestBase struct {
	driver             selenium.WebDriver
	service            *selenium.Service
	verificationErrors []string
}

func (b *TestBase) SetUp(t *testing.T) {
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
	wd.MaximizeWindow("")

	b.driver = wd
	b.service = service
	b.verificationErrors = []string{}
}

func (b *TestBase) TearDown(t *testing.T) {
	//b.driver.Quit()
	//b.service.Stop()
	//if len(b.verificationErrors) > 0 {
	//	t.Errorf("verification errors: %v", b.verificationErrors)
	//}
}

// — вспомогательные методы, вынесенные из теста —

func (b *TestBase) GoToLoginPage() {
	b.driver.Get(baseURL + "/login")
}

func (b *TestBase) GoToMainPage() {
	b.driver.Get(baseURL + "/")
}

func (b *TestBase) Login(user AccountData) {
	username, _ := b.driver.FindElement(selenium.ByName, "username")
	username.Clear()
	username.SendKeys(user.Username)

	password, _ := b.driver.FindElement(selenium.ByName, "password")
	password.Clear()
	password.SendKeys(user.Password)

	btn, _ := b.driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	btn.Click()
}

func (b *TestBase) CreateTodo(todo TodoData) {
	title, _ := b.driver.FindElement(selenium.ByName, "title")
	title.Clear()
	title.SendKeys(todo.Title)

	deadline, _ := b.driver.FindElement(selenium.ByName, "deadline")
	deadline.Clear()
	deadline.SendKeys(todo.Deadline)

	btn, _ := b.driver.FindElement(selenium.ByXPATH, "//h2[text()='Создать задачу']/following-sibling::form//button")
	btn.Click()
}

func (b *TestBase) IsElementPresent(by, value string) bool {
	_, err := b.driver.FindElement(by, value)
	return err == nil
}
