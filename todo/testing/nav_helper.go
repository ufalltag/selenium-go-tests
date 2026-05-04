package seleniumtests

type NavigationHelper struct {
	HelperBase
}

func (n *NavigationHelper) GoToLoginPage() error {
	return n.driver.Get(baseURL + "/login")
}

func (n *NavigationHelper) GoToRegisterPage() error {
	return n.driver.Get(baseURL + "/register")
}

func (n *NavigationHelper) GoToMainPage() error {
	return n.driver.Get(baseURL + "/")
}

func (n *NavigationHelper) CurrentURL() (string, error) {
	return n.driver.CurrentURL()
}
