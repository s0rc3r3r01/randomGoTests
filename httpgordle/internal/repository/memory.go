// Package repository
package repository

import (
	"fmt"
	"log"
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

// Find a game based on its ID. If nothing is found, return a nil pointer and an ErrNotFound error.
func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	log.Printf("Looking for game %s...", id)

	game, found := gr.storage[id]
	if !found {
		return session.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

// Update a game in the database, overwriting it.
func (gr *GameRepository) Update(game session.Game) error {
	_, found := gr.storage[game.ID]
	if !found {
		return fmt.Errorf("can't find game %s: %w", game.ID, ErrNotFound)
	}

	gr.storage[game.ID] = game
	return nil
}
