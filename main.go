package main

import (
	"github.com/medmes/go-tdd/fx/httphandler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	httphandler.New(mux)

	http.ListenAndServe(":8080", mux)
}
