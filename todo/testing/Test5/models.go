package test5

type AccountData struct {
	Username string
	Password string
}

// TodoData — XML-теги нужны для сериализации/десериализации (Инструкция 5)
type TodoData struct {
	Title    string `xml:"Title"`
	Deadline string `xml:"Deadline"`
}
