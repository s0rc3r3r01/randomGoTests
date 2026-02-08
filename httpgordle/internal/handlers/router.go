package handlers

import (
	"net/http"
	"randomGoTests/httpgordle/internal/api"
	"randomGoTests/httpgordle/internal/handlers/newgame"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}
