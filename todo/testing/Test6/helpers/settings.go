package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppSettings struct {
	BaseURL  string `json:"baseUrl"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Settings — глобальная конфигурация, аналог статического класса Settings в C#.
// Инициализируется один раз перед первым использованием (аналог static-конструктора).
var Settings *AppSettings

func init() {
	const file = "settings.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(fmt.Sprintf("settings file not found: %s", file))
	}
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	Settings = &AppSettings{}
	if err := json.Unmarshal(data, Settings); err != nil {
		panic(err)
	}
}
