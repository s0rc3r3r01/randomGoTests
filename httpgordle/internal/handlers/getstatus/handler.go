// Package getstatus
package getstatus

import (
	"encoding/json"
	"log"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/session"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of teh game", http.StatusBadRequest)
		return
	}
	log.Printf("retrieve statuts of game with id: %v", id)

	game := getGame(id)
	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write JSON response: %s", err)
	}

}

func getGame(id string) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
