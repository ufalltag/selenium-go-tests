package helpers

type NavigationHelper struct {
	HelperBase
}

// OpenHomePage — открывает домашнюю страницу (используется в TestBase.Setup).
func (n *NavigationHelper) OpenHomePage() {
	n.driver.Get(Settings.BaseURL)
}

func (n *NavigationHelper) GoToLoginPage() {
	n.driver.Get(Settings.BaseURL + "/login")
}

func (n *NavigationHelper) GoToMainPage() {
	n.driver.Get(Settings.BaseURL + "/")
}

func (n *NavigationHelper) CurrentURL() (string, error) {
	return n.driver.CurrentURL()
}

// IsOnMainPage ждёт появления элемента главной страницы через implicit wait.
func (n *NavigationHelper) IsOnMainPage() bool {
	_, err := n.driver.FindElement("xpath", "//h2[text()='Создать задачу']")
	return err == nil
}
