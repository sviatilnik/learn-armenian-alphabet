package main

import (
	"bufio"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/game"
	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/statistic"
	"github.com/sviatilnik/learn-armenian-alphabet.git/internal/ui"
)

func main() {
	w := os.Stdout
	r := os.Stdin

	ui.Title(w)
	ui.Menu(w)

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
					ui.InvalidMode(w)
				}

				g.Play(ctx, w, r)
			}
		}
	}()

	<-ctx.Done()

	if pGame, ok := g.(statistic.PrinterTracker); ok {
		pGame.PrintStats(w)
	}
	ui.End(w)
}
