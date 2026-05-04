package seleniumtests

import "github.com/tebeka/selenium"

// HelperBase — базовый класс для всех хелперов (Инструкция 3: HelperBase)
// Хранит ссылку на менеджер и драйвер, чтобы не объявлять их в каждом хелпере.
type HelperBase struct {
	manager *AppManager
	driver  selenium.WebDriver
}

func newHelperBase(manager *AppManager) HelperBase {
	return HelperBase{
		manager: manager,
		driver:  manager.driver,
	}
}
