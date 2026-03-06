package game

import (
	"context"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Mode int

const (
	Letter   Mode = 1
	Sound    Mode = 2
	Alphabet Mode = 3 // печать всего алфавита, без статистики
	Unknown  Mode = 0
)

var ErrUnknownGameMode = errors.New("unknown game mode")

func GetModeByString(mode string) Mode {
	if strings.TrimSpace(mode) == "" {
		return Unknown
	}

	convMode, err := strconv.Atoi(mode)
	if err != nil {
		return Unknown
	}

	switch convMode {
	case 1:
		return Letter
	case 2:
		return Sound
	case 3:
		return Alphabet
	default:
		return Unknown
	}
}

type Game interface {
	Play(ctx context.Context, w io.Writer, r io.Reader)
}

func NewGameWithMode(mode Mode) (Game, error) {
	switch mode {
	case Letter:
		return new(LetterGame), nil
	case Sound:
		return new(SoundGame), nil
	case Alphabet:
		return new(AlphabetGame), nil
	case Unknown:
		return nil, ErrUnknownGameMode
	default:
		return nil, ErrUnknownGameMode
	}
}
