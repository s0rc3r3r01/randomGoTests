// Package guess
package guess

import (
	"encoding/json"
	"log"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/repository"
	"randomGoTests/httpgordle/internal/session"
)

func Handler(db *repository.GameRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r := api.GuessRequest{}
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(req.Body).Decode(&r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game, err := guess(id, r, db)
		if err != nil {
			log.Printf("failed to decode request")
			http.Error(w, "broken logic", http.StatusBadRequest)
		}

		apiGame := api.ToGameResponse(game)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}
func guess(id string, r api.GuessRequest, db *repository.GameRepository) (session.Game, error) {
	return session.Game{
		ID: session.GameID(id),
	}, nil
}
