package helpers

import (
	"fmt"
	"strings"
	"todo/test6/models"

	"github.com/tebeka/selenium"
)

type TodoHelper struct {
	HelperBase
}

// TodoItem — данные задачи как они отображаются на странице.
// Deadline в формате DD.MM.YYYY (отображаемый), ID — числовой идентификатор из БД.
type TodoItem struct {
	ID       string
	Title    string
	Deadline string // DD.MM.YYYY
}

func (t *TodoHelper) CreateTodo(todo models.TodoData) {
	title, _ := t.driver.FindElement(selenium.ByName, "title")
	title.Clear()
	title.SendKeys(todo.Title)

	// SendKeys на <input type="date"> ненадёжен на macOS Chrome — браузер
	// интерпретирует символы как нажатия клавиш по секциям (MM/DD/YYYY),
	// а не как строку. JavaScript устанавливает value напрямую.
	t.driver.ExecuteScript(
		`document.querySelector('input[name="deadline"]').value = arguments[0]`,
		[]interface{}{todo.Deadline},
	)

	btn, _ := t.driver.FindElement(selenium.ByXPATH, "//h2[text()='Создать задачу']/following-sibling::form//button")
	btn.Click()
}

// FindTodo ищет задачу по заголовку и возвращает все её поля со страницы.
// Аналог метода, возвращающего конкретный объект для assertEqual-проверок.
func (t *TodoHelper) FindTodo(title string) (TodoItem, bool) {
	xpath := fmt.Sprintf("//div[strong[text()='%s']]", title)
	div, err := t.driver.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return TodoItem{}, false
	}

	// Извлекаем дедлайн из текста дива: "Заголовок — до 01.05.2026\nУдалить"
	text, _ := div.Text()
	deadline := ""
	if parts := strings.SplitN(text, " — до ", 2); len(parts) == 2 {
		deadline = strings.Fields(parts[1])[0]
	}

	// Извлекаем ID из action формы.
	// GetAttribute("action") возвращает полный URL: http://localhost:8080/todos/123/delete
	// Поэтому берём второй элемент с конца, а не с начала.
	form, _ := div.FindElement(selenium.ByXPATH, ".//form")
	action, _ := form.GetAttribute("action")
	id := ""
	if parts := strings.Split(action, "/"); len(parts) >= 2 {
		id = parts[len(parts)-2]
	}

	return TodoItem{ID: id, Title: title, Deadline: deadline}, true
}

// DeleteTodo удаляет задачу по её ID из БД — точнее чем поиск по заголовку.
func (t *TodoHelper) DeleteTodo(id string) {
	// contains нужен потому что action содержит полный URL (http://localhost:8080/todos/123/delete)
	xpath := fmt.Sprintf("//form[contains(@action,'/todos/%s/delete')]//button", id)
	btn, _ := t.driver.FindElement(selenium.ByXPATH, xpath)
	btn.Click()
}

// IsTodoPresent оставлен для обратной совместимости.
func (t *TodoHelper) IsTodoPresent(title string) bool {
	_, ok := t.FindTodo(title)
	return ok
}
