package main

import (
	// "database/sql"
	// "fmt"
	"fmt"
	"log"
	"os"

	"example.com/go-pontifex-tgbot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// var db *sql.DB

// func initDB() {
// 	connStr := "user=botuser password=botpass dbname=mybotdb sslmode=disable"
// 	var err error
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 	}

// 	// Проверка соединения
// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	fmt.Println("Connected to PostgreSQL successfully.")
// }

func safe(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	fn()
}

func main() {
	// initDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	h := handlers.NewUpdateHandler()

	for update := range updates {
		// Если обновление не содержит сообщение, пропускаем его
		if update.Message == nil {
			continue
		}
		go safe(func() {
			h.HandleUpdate(bot, update)
		})
	}
}
