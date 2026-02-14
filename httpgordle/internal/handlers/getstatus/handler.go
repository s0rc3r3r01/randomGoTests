// Package getstatus
package getstatus

import (
	"encoding/json"
	"log"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of teh game", http.StatusBadRequest)
		return
	}
	log.Printf("retrieve statuts of game with id: %v", id)
	apiGame := api.GameResponse{
		ID: id,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write JSON response: %s", err)
	}

}
