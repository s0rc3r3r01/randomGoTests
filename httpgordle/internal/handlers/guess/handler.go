// Package guess
package guess

import (
	"encoding/json"
	"net/http"
	"randomGoTests/httpgordle/internal/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	r := api.GuessRequest{}
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
