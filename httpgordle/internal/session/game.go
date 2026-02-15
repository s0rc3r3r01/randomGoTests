// Package session
package session

import "errors"

type GameID string
type Status string

type Guess struct {
	Word     string
	Feedback string
}

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

type Game struct {
	ID           GameID
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

var ErrGameOver = errors.New("Game over")
