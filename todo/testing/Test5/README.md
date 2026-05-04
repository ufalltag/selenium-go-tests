# Test5 — Инструкция 5: Генератор тестовых данных + XML

## Что требовалось по ТЗ

- Написать консольный генератор тестовых данных
- Генератор принимает параметры: действие, количество, имя файла, формат
- Вывести тестовые данные в XML-файл
- В тестах загружать данные из XML (аналог `TestCaseSource` в C#/NUnit)
- Каждый объект из XML запускается как отдельный тест

## Что сделано

### Файлы

| Файл | Описание |
|---|---|
| `models.go` | `AccountData`, `TodoData` с XML-тегами (`xml:"Title"` и др.) |
| `helper_base.go` | `HelperBase` |
| `app_manager.go` | `AppManager` синглтон |
| `nav_helper.go` | `NavigationHelper` |
| `auth_helper.go` | `AuthHelper` |
| `todo_helper.go` | `TodoHelper` |
| `main_test.go` | `TestMain` |
| `todo_xml_test.go` | Параметризованный тест — читает задачи из `todos.xml` |
| `cmd/generator/main.go` | Консольный генератор тестовых данных |

### Генератор

```bash
# Синтаксис
go run ./cmd/generator <действие> <количество> <файл> <формат>

# Пример — создать 3 задачи в todos.xml
go run ./cmd/generator g 3 todos.xml xml
```

Пример сгенерированного `todos.xml`:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<ArrayOfTodoData>
  <TodoData>
    <Title>Написать тесты 1</Title>
    <Deadline>2026-05-10</Deadline>
  </TodoData>
  <TodoData>
    <Title>Деплой в прод 2</Title>
    <Deadline>2026-05-18</Deadline>
  </TodoData>
</ArrayOfTodoData>
```

### Параметризованный тест (аналог TestCaseSource)

Каждая задача из XML запускается как отдельный подтест через `t.Run`:

```
TestCreateTodoFromXML
  ├── TestCreateTodoFromXML/Написать_тесты_1     — PASS
  ├── TestCreateTodoFromXML/Деплой_в_прод_2      — PASS
  └── TestCreateTodoFromXML/Ревью_кода_3         — PASS
```

## Запуск

```bash
# Шаг 1 — сгенерировать данные
go run ./cmd/generator g 3 todos.xml xml

# Шаг 2 — запустить тесты
go test -v -count=1 ./...
```
