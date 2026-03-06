package internal

import "math/rand/v2"

type Alphabet struct {
}

func (a *Alphabet) RandomLetter() Letter {
	return letters[rand.IntN(len(letters))]
}

type Letter struct {
	Uppercase rune   // Заглавная форма (Ա, Բ, Գ...)
	Lowercase rune   // Строчная форма (ա, բ, գ...)
	Name      string // Название буквы (напр. "айб", "бен", "гим")
	Sound     string // Звук (напр. "a", "b", "g")
}

var letters = []Letter{
	{
		Uppercase: 'Ա',
		Lowercase: 'ա',
		Name:      "айб",
		Sound:     "а",
	},
	{
		Uppercase: 'Բ',
		Lowercase: 'բ',
		Name:      "бен",
		Sound:     "б",
	},
	{
		Uppercase: 'Գ',
		Lowercase: 'գ',
		Name:      "гим",
		Sound:     "г",
	},
	{
		Uppercase: 'Դ',
		Lowercase: 'դ',
		Name:      "да",
		Sound:     "д",
	},
	{
		Uppercase: 'Ե',
		Lowercase: 'ե',
		Name:      "едж",
		Sound:     "е",
	},
	{
		Uppercase: 'Զ',
		Lowercase: 'զ',
		Name:      "за",
		Sound:     "з",
	},
	{
		Uppercase: 'Է',
		Lowercase: 'է',
		Name:      "э",
		Sound:     "э",
	},
	{
		Uppercase: 'Ը',
		Lowercase: 'ը',
		Name:      "ыт",
		Sound:     "э",
	},
	{
		Uppercase: 'Թ',
		Lowercase: 'թ',
		Name:      "то",
		Sound:     "т",
	},
	{
		Uppercase: 'Ժ',
		Lowercase: 'ժ',
		Name:      "же",
		Sound:     "ж",
	},
	{
		Uppercase: 'Ի',
		Lowercase: 'ի',
		Name:      "ини",
		Sound:     "и",
	},
	{
		Uppercase: 'Լ',
		Lowercase: 'լ',
		Name:      "льюн",
		Sound:     "л",
	},
	{
		Uppercase: 'Խ',
		Lowercase: 'խ',
		Name:      "хе",
		Sound:     "х",
	},
	{
		Uppercase: 'Ծ',
		Lowercase: 'ծ',
		Name:      "тьца",
		Sound:     "ц",
	},
	{
		Uppercase: 'Կ',
		Lowercase: 'կ',
		Name:      "кен",
		Sound:     "к",
	},
	{
		Uppercase: 'Հ',
		Lowercase: 'հ',
		Name:      "хо",
		Sound:     "х",
	},
	{
		Uppercase: 'Ձ',
		Lowercase: 'ձ',
		Name:      "дза",
		Sound:     "дз",
	},
	{
		Uppercase: 'Ղ',
		Lowercase: 'ղ',
		Name:      "гхат",
		Sound:     "г",
	},
	{
		Uppercase: 'Ճ',
		Lowercase: 'ճ',
		Name:      "тче",
		Sound:     "тч",
	},
	{
		Uppercase: 'Մ',
		Lowercase: 'մ',
		Name:      "мен",
		Sound:     "м",
	},
	{
		Uppercase: 'Յ',
		Lowercase: 'յ',
		Name:      "йи",
		Sound:     "у",
	},
	{
		Uppercase: 'Ն',
		Lowercase: 'ն',
		Name:      "ну",
		Sound:     "н",
	},
	{
		Uppercase: 'Շ',
		Lowercase: 'շ',
		Name:      "ша",
		Sound:     "ш",
	},
	{
		Uppercase: 'Ո',
		Lowercase: 'ո',
		Name:      "во",
		Sound:     "о",
	},
	{
		Uppercase: 'Չ',
		Lowercase: 'չ',
		Name:      "ча",
		Sound:     "ч",
	},
	{
		Uppercase: 'Պ',
		Lowercase: 'պ',
		Name:      "пэ",
		Sound:     "п",
	},
	{
		Uppercase: 'Ջ',
		Lowercase: 'ջ',
		Name:      "дже",
		Sound:     "дж",
	},
	{
		Uppercase: 'Ռ',
		Lowercase: 'ռ',
		Name:      "ра",
		Sound:     "р",
	},
	{
		Uppercase: 'Ս',
		Lowercase: 'ս',
		Name:      "се",
		Sound:     "с",
	},
	{
		Uppercase: 'Վ',
		Lowercase: 'վ',
		Name:      "вев",
		Sound:     "в",
	},
	{
		Uppercase: 'Տ',
		Lowercase: 'տ',
		Name:      "тюн",
		Sound:     "т",
	},
	{
		Uppercase: 'Ր',
		Lowercase: 'ր',
		Name:      "ре",
		Sound:     "р",
	},
	{
		Uppercase: 'Ց',
		Lowercase: 'ց',
		Name:      "цо",
		Sound:     "ц",
	},
	{
		Uppercase: 'Ց',
		Lowercase: 'ց',
		Name:      "цо",
		Sound:     "ц",
	},
	{
		Uppercase: 'Ւ',
		Lowercase: 'ւ',
		Name:      "йун",
		Sound:     "у",
	},
	{
		Uppercase: 'Փ',
		Lowercase: 'փ',
		Name:      "пьюр",
		Sound:     "п",
	},
	{
		Uppercase: 'Ք',
		Lowercase: 'ք',
		Name:      "ке",
		Sound:     "к",
	},
	{
		Uppercase: 'և',
		Lowercase: 'և',
		Name:      "ев",
		Sound:     "ев",
	},
	{
		Uppercase: 'Օ',
		Lowercase: 'օ',
		Name:      "о",
		Sound:     "о",
	},
	{
		Uppercase: 'Ֆ',
		Lowercase: 'ֆ',
		Name:      "фе",
		Sound:     "ф",
	},
}
