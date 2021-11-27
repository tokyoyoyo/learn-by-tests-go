package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
