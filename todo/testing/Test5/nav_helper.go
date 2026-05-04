package test5

type NavigationHelper struct {
	HelperBase
}

func (n *NavigationHelper) GoToLoginPage() {
	n.driver.Get(baseURL + "/login")
}
