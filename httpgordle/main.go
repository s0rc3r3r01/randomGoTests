package main

import (
	"fmt"
	"net/http"
	"os"
	"randomGoTests/httpgordle/internal/handlers"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
