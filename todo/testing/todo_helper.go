package seleniumtests

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type TodoHelper struct {
	HelperBase
}

func (t *TodoHelper) CreateTodo(todo TodoData) error {
	titleField, err := t.driver.FindElement(selenium.ByName, "title")
	if err != nil {
		return err
	}
	titleField.Clear()
	titleField.SendKeys(todo.Title)

	deadlineField, err := t.driver.FindElement(selenium.ByName, "deadline")
	if err != nil {
		return err
	}
	deadlineField.Clear()
	deadlineField.SendKeys(todo.Deadline)

	submitBtn, err := t.driver.FindElement(selenium.ByXPATH, "//h2[text()='Создать задачу']/following-sibling::form//button[@type='submit']")
	if err != nil {
		return err
	}
	return submitBtn.Click()
}

// DeleteTodoByTitle — удаляет задачу по названию (Инструкция 4: тест удаления)
func (t *TodoHelper) DeleteTodoByTitle(title string) error {
	xpath := fmt.Sprintf("//div[strong[text()='%s']]//button[@type='submit']", title)
	btn, err := t.driver.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return err
	}
	return btn.Click()
}

func (t *TodoHelper) GetTodoTitles() ([]string, error) {
	elements, err := t.driver.FindElements(selenium.ByXPATH, "//div/strong")
	if err != nil {
		return nil, err
	}
	titles := make([]string, 0, len(elements))
	for _, e := range elements {
		text, err := e.Text()
		if err != nil {
			return nil, err
		}
		titles = append(titles, text)
	}
	return titles, nil
}

func (t *TodoHelper) IsTodoPresent(title string) (bool, error) {
	titles, err := t.GetTodoTitles()
	if err != nil {
		return false, err
	}
	for _, s := range titles {
		if s == title {
			return true, nil
		}
	}
	return false, nil
}
