// Генератор тестовых данных (Инструкция 5)
// Запуск: go run ./cmd/generator g <кол-во> <файл> xml
// Пример: go run ./cmd/generator g 3 todos.xml xml
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var todoTitles = []string{
	"Написать отчёт",
	"Провести встречу",
	"Обновить документацию",
	"Исправить баг",
	"Ревью кода",
	"Написать тесты",
	"Деплой в прод",
	"Настроить CI/CD",
	"Подготовить презентацию",
	"Провести ретроспективу",
}

type TodoData struct {
	Title    string `xml:"Title"`
	Deadline string `xml:"Deadline"`
}

type ArrayOfTodoData struct {
	XMLName xml.Name   `xml:"ArrayOfTodoData"`
	Items   []TodoData `xml:"TodoData"`
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: generator <action> <count> <filename> <format>")
		fmt.Println("Example: generator g 3 todos.xml xml")
		os.Exit(1)
	}

	action := os.Args[1]
	if action != "g" {
		log.Fatalf("Unknown action %q — use 'g' to generate", action)
	}

	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count <= 0 {
		log.Fatalf("Invalid count: %s", os.Args[2])
	}

	filename := os.Args[3]
	format := os.Args[4]

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	todos := make([]TodoData, count)
	for i := range todos {
		todos[i] = TodoData{
			Title:    fmt.Sprintf("%s %d", todoTitles[rng.Intn(len(todoTitles))], i+1),
			Deadline: time.Now().AddDate(0, 0, rng.Intn(30)+1).Format("2006-01-02"),
		}
	}

	switch format {
	case "xml":
		writeXML(filename, todos)
	default:
		log.Fatalf("Unknown format %q — use 'xml'", format)
	}

	fmt.Printf("Generated %d todos → %s\n", count, filename)
}

func writeXML(filename string, todos []TodoData) {
	data, err := xml.MarshalIndent(ArrayOfTodoData{Items: todos}, "", "  ")
	if err != nil {
		log.Fatalf("XML marshal: %v", err)
	}
	content := append([]byte(xml.Header), data...)
	if err := os.WriteFile(filename, content, 0644); err != nil {
		log.Fatalf("Write file: %v", err)
	}
}
