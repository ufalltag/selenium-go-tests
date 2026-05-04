package seleniumtests

// AccountData — данные пользователя (Инструкция 2: класс для объекта сайта)
type AccountData struct {
	Username string `xml:"Username"`
	Password string `xml:"Password"`
}

// TodoData — данные задачи (Инструкция 2: класс для объекта сайта)
type TodoData struct {
	Title    string `xml:"Title"`
	Deadline string `xml:"Deadline"`
}
