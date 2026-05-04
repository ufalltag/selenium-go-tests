// Инструкция 3: TestBase знает только об AppManager.
// SetUp/TearDown остались в TestBase, но логика инициализации перенесена в AppManager.
package test3

type TestBase struct {
	app *AppManager
}

func (b *TestBase) SetUp() {
	b.app = NewAppManager()
}

func (b *TestBase) TearDown() {
	b.app.Stop()
}
