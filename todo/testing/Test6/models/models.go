package models

type AccountData struct {
	Username string
	Password string
}

type TodoData struct {
	Title    string `xml:"Title"`
	Deadline string `xml:"Deadline"`
}
