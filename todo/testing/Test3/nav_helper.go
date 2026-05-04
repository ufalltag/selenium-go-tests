package test3

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
