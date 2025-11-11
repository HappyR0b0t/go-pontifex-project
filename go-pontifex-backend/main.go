package main

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	"example.com/go-pontifex/handlers"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	// connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	user, password, host, port, dbname,
	// )

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/cipher", handlers.CipherHandler)
	http.HandleFunc("/decipher", handlers.DecipherHandler)
	http.HandleFunc("/generate", handlers.GenerateDeckHandler)

	log.Info().Msg("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(errors.New("server failed to start"))
	}
}
