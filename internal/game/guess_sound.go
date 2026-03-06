package game

import (
	"context"
	"io"
)

type SoundGame struct {
}

func (g *SoundGame) Play(ctx context.Context, w io.Writer, r io.Reader) {
	w.Write([]byte("Режим в разработке"))
}
