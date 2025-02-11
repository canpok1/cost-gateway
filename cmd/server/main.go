package main

import (
	"log"
	"net/http"

	"github.com/canpok1/code-gateway/internal/api"
)

func main() {
	server := api.NewServer()

	r := http.NewServeMux()

	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
