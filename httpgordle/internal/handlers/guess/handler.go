// Package guess
package guess

import (
	"encoding/json"
	"log"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
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
	apiGame := api.GameResponse{
		ID: id,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}
