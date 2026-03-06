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
	printEnd(w)

	if pGame, ok := g.(statistic.PrinterTracker); ok {
		pGame.PrintStats(w)
	}
}

func printTitle(w io.Writer) {
	w.Write([]byte("Добро пожаловать в тренажёр армянского алфавита!\n"))
}

func printMenu(w io.Writer) {
	w.Write([]byte(`Выберите режим:
[1] Назвать букву (по символу)
[2] Назвать звук/транслитерацию (по символу)
Выбор:`))
}

func printEnd(w io.Writer) {
	w.Write([]byte("\n"))
	w.Write([]byte("Спасибо за игру!\n"))
}
