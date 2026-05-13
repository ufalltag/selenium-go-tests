// Генератор тестовых данных.
// Запуск: go run ./cmd/generator g <кол-во> <файл> xml
// Пример: go run ./cmd/generator g 3 todos.xml xml
package main

import (
	cryptorand "crypto/rand"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var titles = []string{
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
		fmt.Fprintln(os.Stderr, "Usage: generator g <count> <filename> xml")
		fmt.Fprintln(os.Stderr, "Example: generator g 3 todos.xml xml")
		os.Exit(1)
	}

	action := os.Args[1]
	if action != "g" {
		log.Fatalf("Unknown action %q — use 'g'", action)
	}

	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count <= 0 {
		log.Fatalf("Invalid count: %s", os.Args[2])
	}

	filename := os.Args[3]
	format := os.Args[4]
	if format != "xml" {
		log.Fatalf("Unknown format %q — use 'xml'", format)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// batchID делает заголовки уникальными между запусками генератора,
	// чтобы FindTodo в тестах находил именно свежесозданную задачу.
	b := make([]byte, 3)
	cryptorand.Read(b)
	batchID := fmt.Sprintf("%x", b)

	items := make([]TodoData, count)
	for i := range items {
		items[i] = TodoData{
			Title:    fmt.Sprintf("%s %d [%s]", titles[rng.Intn(len(titles))], i+1, batchID),
			Deadline: time.Now().AddDate(0, 0, rng.Intn(30)+1).Format("2006-01-02"),
		}
	}

	out, err := xml.MarshalIndent(ArrayOfTodoData{Items: items}, "", "  ")
	if err != nil {
		log.Fatalf("xml marshal: %v", err)
	}

	content := append([]byte(xml.Header), out...)
	if err := os.WriteFile(filename, content, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	fmt.Printf("Generated %d todos → %s\n", count, filename)
}
