package main

import (
	// "database/sql"
	// "fmt"
	"log"
	"os"
	"strings"
	"sync"

	handlers "example.com/go-pontifex-tgbot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	userStates        = make(map[int64]string)
	userStatesMu      sync.Mutex
	textForDecipher   = make(map[int64]string)
	textForDecipherMu sync.Mutex

	// Command texts
	cipherText   = "Please, enter message to cipher. Message should contain only latin characters, no symbols allowed"
	decipherText = "Please, enter ciphered message. Message should contain only latin characters, no symbols allowed"
	helpText     = `For this algorithm, to be able to cipher a text message, it should consist of latin characters only 
	(a-z), no symbols! The case of characters is irrelevant. Ciphered and deciphered message will be a string consisting 
	of upper case characters (A-Z) grouped in 5 letters, for example:
    "Do not use pc" after ciphering will turn into a random set of characters that could look like this - "ZFQIK LCOUA"
    "ZFQIK LCOUA" after deciphering will turn into "DONOT USEPC". You get the idea. This grouping by 5 characters is just 
	a common convention in cryptography.
	It is important to mention, that in order to sucssesfully decipher a message, the same deck should be used. 
	The deck is a text containing 54 "words" of "cards" in following format <suit>-<rank>. Suits are classical: clubs, 
	diamond, hearts and spades. Ranks are as follows: Ace, 2, 3 ... King and two Jokers: JA and JB. If you want another 
	person to be able to decipher your message, you should provide a ciphered text and the deck you used to cipher it (in its starting state).`
	aboutText = `Often interpreted as a compound originally meaning bridge\n 
	maker from Proto Italic "pontifaks" equivalent to "pons" bridge\n
	"fex" suffix representing a maker or producer either metaphorically\n 
	one who negotiates between gods and men or literally if at some\n
	point the social class which supplied the priests was more or less\n 
	identical with engineers that were responsible for building bridges\n
	Pontifex is an algorithm for ciphering and deciphering text messages, 
	described in Neal Stephenson's novel "Cryptonomicon", created by Bruce 
	Schneier. More widely known as Solitaire cipher, because it uses a deck 
	of playing cards. Deck consits of 52 classic cards, plus two Jokers. 
	This ciphering algorithm can be performed with physical deck of cards, 
	but I decided to make a program for study puprposes and fun.`

	// Menu texts
	firstMenu  = "<b>Main menu</b>\n\nSelect an option"
	secondMenu = "<b>Cipher a text message menu</b>\n\nSelect an option"
	thirdMenu  = "<b>Decipher a text message menu</b>\n\nSelect an option"
	fourthMenu = "<b>About</b>\n\nAbout Pontifex algorithm..."

	// Button texts
	aboutButton        = "About"
	backButton         = "Back"
	cipherButton       = "Cipher text"
	decipherButton     = "Decipher text"
	loadtextButton     = "Load text"
	loaddeckButton     = "Load deck"
	generatedeckButton = "Generate deck"
	// tutorialButton     = "Tutorial"

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(aboutButton, aboutButton),
			tgbotapi.NewInlineKeyboardButtonData(cipherButton, cipherButton),
			tgbotapi.NewInlineKeyboardButtonData(decipherButton, decipherButton),
			tgbotapi.NewInlineKeyboardButtonData(generatedeckButton, generatedeckButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
			tgbotapi.NewInlineKeyboardButtonData(loadtextButton, loadtextButton),
			tgbotapi.NewInlineKeyboardButtonData(loaddeckButton, loaddeckButton),
			tgbotapi.NewInlineKeyboardButtonData(generatedeckButton, generatedeckButton),
		),
	)
	// Keyboard layout for the second menu. Two buttons, one per row
	thirdMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
			tgbotapi.NewInlineKeyboardButtonData(loadtextButton, loadtextButton),
			tgbotapi.NewInlineKeyboardButtonData(loaddeckButton, loaddeckButton),
		),
	)
	fourthMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
			// tgbotapi.NewInlineKeyboardButtonData(cipherButton, cipherButton),
		),
	)
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

	for update := range updates {
		// Если обновление не содержит сообщение, пропускаем его
		if update.Message == nil {
			continue
		}
		go handleUpdate(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(bot, update.Message)
	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(bot, update.CallbackQuery)
	}
}

func stateChanger(chatID int64, state string) {
	userStatesMu.Lock()
	userStates[chatID] = state
	userStatesMu.Unlock()
}

func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	userStatesMu.Lock()
	state, ok := userStates[chatID]
	if !ok {
		state = "started"
		userStates[chatID] = state
	}
	userStatesMu.Unlock()

	if msg.IsCommand() {
		switch msg.Command() {
		case "cipher":
			reply := tgbotapi.NewMessage(chatID, cipherText)
			bot.Send(reply)
			stateChanger(chatID, "waiting_for_text_to_cipher")
			return

		case "decipher":
			reply := tgbotapi.NewMessage(chatID, decipherText)
			bot.Send(reply)
			stateChanger(chatID, "waiting_for_text_to_decipher")
			return

		case "generate":
			replyMessage := tgbotapi.NewMessage(chatID, "Here is your deck:")
			bot.Send(replyMessage)
			generated := handlers.HandleGenerateCommand()
			reply := tgbotapi.NewMessage(chatID, generated)
			bot.Send(reply)
			return

		case "start":
			reply := tgbotapi.NewMessage(chatID, "Use /cipher or /decipher to start.")
			bot.Send(reply)
			stateChanger(chatID, "started")
			return

		case "menu":
			sendMenu(bot, chatID)
			return

		case "help":
			reply := tgbotapi.NewMessage(chatID, helpText)
			bot.Send(reply)
			return

		case "about":
			reply := tgbotapi.NewMessage(chatID, aboutText)
			bot.Send(reply)
			return

		default:
			reply := tgbotapi.NewMessage(chatID, "Sorry, no such command. Try again")
			bot.Send(reply)
			return
		}

	}

	if !msg.IsCommand() && msg.Text != "" {
		switch state {
		case "started":
			reply := tgbotapi.NewMessage(chatID, "Please use /cipher, /generate or /decipher first.")
			bot.Send(reply)

		case "waiting_for_text_to_cipher":
			textForDecipherMu.Lock()
			textForDecipher[chatID] = msg.Text
			textForDecipherMu.Unlock()

			reply := tgbotapi.NewMessage(chatID, "Provide a deck or send 'no deck' message")
			bot.Send(reply)

			stateChanger(chatID, "waiting_for_deck_to_cipher")

		case "waiting_for_deck_to_cipher":
			textForDecipherMu.Lock()
			originalText := textForDecipher[chatID]
			textForDecipherMu.Unlock()

			cipheredTextMessage := tgbotapi.NewMessage(chatID, "Here is your ciphered text:")
			bot.Send(cipheredTextMessage)

			cipheredText := handlers.HandleCipherCommand(originalText, msg.Text)
			replyText := tgbotapi.NewMessage(chatID, cipheredText[0])
			bot.Send(replyText)

			DeckMessage := tgbotapi.NewMessage(chatID, "Here is your deck:")
			bot.Send(DeckMessage)

			replyDeck := tgbotapi.NewMessage(chatID, strings.Join(cipheredText[1:], " "))
			bot.Send(replyDeck)

			stateChanger(chatID, "started")

		case "waiting_for_text_to_decipher":
			textForDecipherMu.Lock()
			textForDecipher[chatID] = msg.Text
			textForDecipherMu.Unlock()

			reply := tgbotapi.NewMessage(chatID, "Now send the deck to use for deciphering:")
			bot.Send(reply)

			stateChanger(chatID, "waiting_for_deck_to_decipher")

		case "waiting_for_deck_to_decipher":
			textForDecipherMu.Lock()
			originalText := textForDecipher[chatID]
			textForDecipherMu.Unlock()

			decipheredTextMessage := tgbotapi.NewMessage(chatID, "Here is your deciphered text:")
			bot.Send(decipheredTextMessage)

			decipheredText := handlers.HandleDecipherCommand(originalText, msg.Text)
			replyText := tgbotapi.NewMessage(chatID, decipheredText[0])
			bot.Send(replyText)

			DeckMessage := tgbotapi.NewMessage(chatID, "Here is your deck:")
			bot.Send(DeckMessage)

			replyDeck := tgbotapi.NewMessage(chatID, strings.Join(decipheredText[1:], " "))
			bot.Send(replyDeck)

			stateChanger(chatID, "started")

		}
	}
}

func handleButton(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == cipherButton {
		text = secondMenu
		markup = secondMenuMarkup
	} else if query.Data == decipherButton {
		text = thirdMenu
		markup = thirdMenuMarkup
	} else if query.Data == aboutButton {
		text = fourthMenu
		markup = fourthMenuMarkup
	} else if query.Data == backButton {
		text = firstMenu
		markup = firstMenuMarkup
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func sendMenu(bot *tgbotapi.BotAPI, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}
