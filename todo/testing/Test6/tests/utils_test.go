package tests

import (
	cryptorand "crypto/rand"
	"fmt"
	"strings"
	"testing"
)

// assertEqualStr — аналог assertEqual из C#/NUnit.
func assertEqualStr(t *testing.T, field, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("%s: want %q, got %q", field, want, got)
	}
}

// newID генерирует короткий уникальный идентификатор для изоляции тестовых данных.
// Без него одинаковые заголовки между запусками мешают assertEqual-проверкам.
func newID() string {
	b := make([]byte, 4)
	cryptorand.Read(b)
	return fmt.Sprintf("%x", b)
}

// toDisplayDate конвертирует YYYY-MM-DD → DD.MM.YYYY для сравнения
// с датой, которую приложение отображает на странице.
func toDisplayDate(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return s
	}
	return parts[2] + "." + parts[1] + "." + parts[0]
}
