package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/game"
	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/statistic"
)

func main() {
	w := os.Stdout
	r := os.Stdin

	printTitle(w)
	printMenu(w)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var g game.Game

	go func() {
		scanner := bufio.NewScanner(r)

		var err error

		for g == nil {
			if scanner.Scan() {
				mode := game.GetModeByString(scanner.Text())

				g, err = game.NewGameWithMode(mode)
				if err != nil {
					w.Write([]byte("Некорректный режим!\n"))
					printMenu(w)
				}

				g.Play(ctx, w, r)
			}
		}
	}()

	<-ctx.Done()

	if pGame, ok := g.(statistic.PrinterTracker); ok {
		pGame.PrintStats(w)
	}
}

func printTitle(w io.Writer) {
	w.Write([]byte("Добро пожаловать в тренажёр армянского алфавита!\n\n"))
	w.Write([]byte("Тренажёр помогает запомнить буквы и их названия/звуки. Выберите режим, отвечайте на вопросы. В любой момент выйти из игры можно по Ctrl+C.\n\n"))
}

func printMenu(w io.Writer) {
	w.Write([]byte(`Выберите режим:
[1] Назвать букву (по символу)
[2] Назвать звук/транслитерацию (по символу)
[3] Показать алфавит
Выбор:`))
}
