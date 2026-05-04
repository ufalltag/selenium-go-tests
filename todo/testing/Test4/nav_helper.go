package test4

type NavigationHelper struct {
	HelperBase
}

func (n *NavigationHelper) GoToLoginPage() {
	n.driver.Get(baseURL + "/login")
}

func (n *NavigationHelper) GoToMainPage() {
	n.driver.Get(baseURL + "/")
}

func (n *NavigationHelper) GoToRegisterPage() {
	n.driver.Get(baseURL + "/register")
}

func (n *NavigationHelper) CurrentURL() (string, error) {
	return n.driver.CurrentURL()
}

// IsOnMainPage ждёт появления элемента главной страницы (implicit wait 10s).
// Надёжнее чем проверка URL сразу после редиректа.
func (n *NavigationHelper) IsOnMainPage() bool {
	_, err := n.driver.FindElement("xpath", "//h2[text()='Создать задачу']")
	return err == nil
}
