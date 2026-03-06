package game

import (
	"context"
	"fmt"
	"io"

	"github.com/sviatilnik/learn-armenian-alphabet.git/internal"
)

// AlphabetGame — режим «Печать алфавита»: выводит все буквы, без статистики.
type AlphabetGame struct {
	alphabet *internal.Alphabet
}

func (g *AlphabetGame) Play(ctx context.Context, w io.Writer, _ io.Reader) {
	if g.alphabet == nil {
		g.alphabet = &internal.Alphabet{}
	}
	for _, l := range g.alphabet.AllLetters() {
		fmt.Fprintf(w, "%c %c — %s, звук: %s\n", l.Uppercase, l.Lowercase, l.Name, l.Sound)
	}
}
