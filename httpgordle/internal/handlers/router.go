package handlers

import (
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/handlers/getstatus"
	"randomGoTests/httpgordle/internal/handlers/newgame"
	"randomGoTests/httpgordle/internal/repository"
)

func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db))
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, getstatus.Handler(db))
	return r
}
