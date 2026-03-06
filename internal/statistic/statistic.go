package statistic

import "io"

type Stat struct {
	Number         int
	DisplayedRune  rune   // Какая буква была показана пользователю
	ExpectedAnswer string // Какой ответ ожидался
	UserAnswer     string // Что ввел пользователь
	IsCorrect      bool   // Результат
}

type Recorder interface {
	Record(stat Stat) error
}

type Tracker interface {
	Record(stat Stat) error
	Stats() ([]Stat, error)
}

type PrinterTracker interface {
	Tracker
	PrintStats(w io.Writer)
}
