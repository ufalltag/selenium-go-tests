// Инструкция 3: базовый класс для всех хелперов.
// Хранит ссылку на менеджер и драйвер — хелперам не нужно объявлять их самим.
package test3

import "github.com/tebeka/selenium"

type HelperBase struct {
	manager *AppManager
	driver  selenium.WebDriver
}
