package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"strings"

	"github.com/sviatilnik/learn-armenian-alphabet.git/internal"
	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/statistic"
)

type LetterGame struct {
	stats    []statistic.Stat
	scan     *bufio.Scanner
	alphabet *internal.Alphabet
}

func (g *LetterGame) Play(ctx context.Context, w io.Writer, r io.Reader) {
	g.scan = bufio.NewScanner(r)
	if g.alphabet == nil {
		g.alphabet = &internal.Alphabet{}
	}

	for turn := 1; ; turn++ {
		if ctx.Err() != nil {
			fmt.Fprintln(w, "Игра завершена.")
			return
		}

		letter := g.alphabet.RandomLetter()
		displayedRune := g.pickCase(letter)
		expectedAnswer := letter.Name

		fmt.Fprintf(w, "Буква: %c\nВаш ответ: ", displayedRune)

		if !g.scan.Scan() {
			break
		}
		userAnswer := strings.TrimSpace(g.scan.Text())

		isCorrect := strings.EqualFold(expectedAnswer, userAnswer)
		if isCorrect {
			fmt.Fprintln(w, "Правильно! 👍")
		} else {
			fmt.Fprintf(w, "Неправильно 👎. Правильный ответ: %s\n", expectedAnswer)
		}

		g.Record(statistic.Stat{
			Number:         turn,
			DisplayedRune:  displayedRune,
			ExpectedAnswer: expectedAnswer,
			UserAnswer:     userAnswer,
			IsCorrect:      isCorrect,
		})
	}
}

func (g *LetterGame) pickCase(l internal.Letter) rune {
	if rand.IntN(2) == 0 {
		return l.Uppercase
	}
	return l.Lowercase
}

func (g *LetterGame) Record(stat statistic.Stat) error {
	g.stats = append(g.stats, stat)
	return nil
}

func (g *LetterGame) Stats() ([]statistic.Stat, error) {
	return g.stats, nil
}

func (g *LetterGame) PrintStats(w io.Writer) {
	if len(g.stats) == 0 {
		return
	}
	total, correct := len(g.stats), 0
	fmt.Fprintln(w, "--------------------------------------------------")
	fmt.Fprintln(w, "Ваши результаты")
	fmt.Fprintln(w, "--------------------------------------------------")
	for _, s := range g.stats {
		verdict := "Верно"
		if !s.IsCorrect {
			verdict = "Неверно"
		} else {
			correct++
		}
		if s.UserAnswer == "" {
			s.UserAnswer = "-"
		}
		fmt.Fprintf(w, "Ход #%d | Буква: %c (ожидание: %s) | Ваш ответ: %s | %s\n",
			s.Number, s.DisplayedRune, s.ExpectedAnswer, s.UserAnswer, verdict)
	}

	fmt.Fprintln(w, "--------------------------------------------------")
	fmt.Fprintf(w, "Всего: %d | Верно: %d  | Неверно: %d | \n", total, correct, total-correct)
	fmt.Fprintln(w, "--------------------------------------------------")
}
