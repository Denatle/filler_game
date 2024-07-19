package logic

import (
	"errors"
	shared "filler/shared"
)

type Logic struct {
	state *shared.GameState
}

func (l Logic) GetAvailableMoves() ([4]int8, error) {
	var moves [4]int8

	return moves, errors.New("Not implemented")
}

func (l Logic) PlayMove(int8) error {
	return errors.New("Not implemented")
}

func (l Logic) Init(state *shared.GameState) {
	l.state = state
}

func (l Logic) Close() {
	return
}
