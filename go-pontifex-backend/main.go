package main

import (
	"log"
	"net/http"

	"example.com/go-pontifex/handlers"
)

func main() {

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/cipher", handlers.CipherHandler)
	http.HandleFunc("/decipher", handlers.DecipherHandler)
	http.HandleFunc("/generate", handlers.GenerateDeckHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
