package main

import (
	"log"
	"net/http"

	"github.com/tokyoyoyo/learn-by-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}
