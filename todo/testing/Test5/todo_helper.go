package test5

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type TodoHelper struct {
	HelperBase
}

func (t *TodoHelper) CreateTodo(todo TodoData) {
	title, _ := t.driver.FindElement(selenium.ByName, "title")
	title.Clear()
	title.SendKeys(todo.Title)

	deadline, _ := t.driver.FindElement(selenium.ByName, "deadline")
	deadline.Clear()
	deadline.SendKeys(todo.Deadline)

	btn, _ := t.driver.FindElement(selenium.ByXPATH, "//h2[text()='Создать задачу']/following-sibling::form//button")
	btn.Click()
}

func (t *TodoHelper) IsTodoPresent(title string) bool {
	xpath := fmt.Sprintf("//div/strong[text()='%s']", title)
	_, err := t.driver.FindElement(selenium.ByXPATH, xpath)
	return err == nil
}
