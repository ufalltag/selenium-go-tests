# Test3 — Инструкция 3: Хелпер-классы и AppManager

## Что требовалось по ТЗ

- Создать хелпер-классы: `NavigationHelper`, `AuthHelper` (LoginHelper), `TodoHelper`
- Создать базовый класс `HelperBase` — хранит драйвер, чтобы не объявлять его в каждом хелпере
- Создать `AppManager` — управляет драйвером и хелперами, хелперы получают менеджера в конструкторе
- `TestBase` теперь знает только об `AppManager`
- Тесты обращаются к хелперам через менеджера: `app.Nav.GoToLoginPage()`

## Что сделано

### Файлы

| Файл | Описание |
|---|---|
| `models.go` | `AccountData`, `TodoData` |
| `helper_base.go` | `HelperBase` — хранит `*AppManager` и `driver` |
| `app_manager.go` | `AppManager` — создаёт драйвер, инициализирует хелперы через `NewAppManager()` |
| `nav_helper.go` | `NavigationHelper` — переходы по страницам |
| `auth_helper.go` | `AuthHelper` — логин, регистрация |
| `todo_helper.go` | `TodoHelper` — создание задачи |
| `test_base.go` | `TestBase` — `SetUp` создаёт `AppManager`, `TearDown` его останавливает |
| `login_test.go` | Тест авторизации |
| `todo_test.go` | Тест создания задачи |

### Схема зависимостей

```
TestBase
  └── AppManager
        ├── NavigationHelper → HelperBase (driver)
        ├── AuthHelper       → HelperBase (driver)
        └── TodoHelper       → HelperBase (driver)
```

### Что намеренно НЕ сделано

- `AppManager` НЕ синглтон — каждый тест открывает свой браузер (это Инструкция 4)
- Нет `Assert`-ов в тестах — это Инструкция 4

## Запуск

> Пользователь `ufalltag` должен быть зарегистрирован заранее на `http://localhost:8080/register`

```bash
go test -v -count=1 ./...
```
