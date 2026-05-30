package main

import (
	"log"
	"net/http"

	"github.com/SakamotoHiroya/go-ogen/internal/gen/greetingapi"
	"github.com/SakamotoHiroya/go-ogen/internal/service"
)

func main() {
	handler := service.NewGreetingService()

	server, err := greetingapi.NewServer(handler)
	if err != nil {
		log.Fatalf("create server: %v", err)
	}

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
