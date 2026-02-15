// Package repository
package repository

import (
	"fmt"
	"randomGoTests/httpgordle/internal/session"
)

type GameRepository struct {
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("gameID %s already exists", game.ID)
	}
	gr.storage[game.ID] = game
	return nil
}
