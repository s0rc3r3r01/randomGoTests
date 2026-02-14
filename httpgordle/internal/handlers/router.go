package handlers

import (
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/handlers/getstatus"
	"randomGoTests/httpgordle/internal/handlers/newgame"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)
	mux.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle)
	return mux
}
