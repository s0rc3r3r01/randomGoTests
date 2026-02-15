// Package getstatus
package getstatus

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
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of teh game", http.StatusBadRequest)
			return
		}
		log.Printf("retrieve statuts of game with id: %v", id)

		game := getGame(id, db)
		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write JSON response: %s", err)
		}

	}
}
func getGame(id string, db *repository.GameRepository) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
