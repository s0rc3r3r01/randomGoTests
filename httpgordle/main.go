package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
}
