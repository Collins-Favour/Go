package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandleFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
