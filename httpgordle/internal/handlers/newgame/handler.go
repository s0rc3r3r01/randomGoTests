// Package newgame is the best
package newgame

import (
	"encoding/json"
	"log"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/session"
)

type GameAdder interface {
	Add(game session.Game) error
}

func Handler(db GameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if game, err := createGame(db); err != nil {
			log.Printf("Unable to create a new game: %s", err)
			http.Error(w, "failed to create new game", http.StatusInternalServerError)
			return
		} else {
			{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				res := response(game)
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					log.Printf("failed to write JSON response: %s", err)
				}
			}

		}
	}
}

func createGame(db GameAdder) (session.Game, error) {
	return session.Game{}, nil
}

func response(game session.Game) api.GameResponse {
	return api.GameResponse{}
}
