# Test4 — Инструкция 4: Один браузер на все тесты + Assert-ы

## Что требовалось по ТЗ

- Сделать `AppManager` синглтоном — все тесты запускаются в одном браузере
- Убрать `TearDown` из каждого теста — браузер закрывается один раз в конце
- Добавить `Assert`-ы после каждого теста для проверки результата
- Добавить новый тест — редактирование или удаление сущности (реализовано удаление)

## Что сделано

### Файлы

| Файл | Описание |
|---|---|
| `models.go` | `AccountData`, `TodoData` |
| `helper_base.go` | `HelperBase` |
| `app_manager.go` | `AppManager` синглтон через `sync.Once` + `GetInstance()` |
| `nav_helper.go` | `NavigationHelper` |
| `auth_helper.go` | `AuthHelper` |
| `todo_helper.go` | `TodoHelper` — добавлены `DeleteTodoByTitle`, `IsTodoPresent` |
| `main_test.go` | `TestMain` — браузер запускается один раз, `app.Stop()` в конце |
| `login_test.go` | Тесты авторизации с Assert-ами |
| `todo_test.go` | `TestCreateTodo` + `TestDeleteTodo` с Assert-ами |

### Ключевые изменения относительно Test3

```
AppManager.GetInstance()  — синглтон вместо NewAppManager()
TestMain                  — заменяет SetUp/TearDown в каждом тесте
Assert-ы                  — t.Errorf / t.Fatal после каждого действия
TestDeleteTodo            — новый тест удаления задачи
```

### Assert-ы в тестах

| Тест | Что проверяется |
|---|---|
| `TestLogin` | После входа URL оканчивается на `/` |
| `TestLoginInvalidCredentials` | После неверных данных URL содержит `/login` |
| `TestCreateTodo` | Задача появилась в списке |
| `TestDeleteTodo` | Задача исчезла из списка после удаления |

## Запуск

```bash
go test -v -count=1 ./...
```
