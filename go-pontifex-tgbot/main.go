package main

import (
	"log"
	"os"
	"strings"
	"sync"

	handlers "example.com/go-pontifex-tgbot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	userStates        = make(map[int64]string)
	userStatesMu      sync.Mutex
	textForDecipher   = make(map[int64]string)
	textForDecipherMu sync.Mutex

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

func main() {
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
			reply := tgbotapi.NewMessage(chatID, "Enter message to encrypt!")
			bot.Send(reply)
			stateChanger(chatID, "waiting_for_text_to_cipher")
			return

		case "decipher":
			reply := tgbotapi.NewMessage(chatID, "Enter message to decrypt!")
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
			reply := tgbotapi.NewMessage(chatID, "Use /cipher or /decipher to start.")
			bot.Send(reply)
			return

		case "about":
			reply := tgbotapi.NewMessage(chatID, "This bot is a simple implementation of the Pontifex algorithm.")
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

			reply := tgbotapi.NewMessage(
				chatID,
				"Now send the deck to use for ciphering.\n If you choose not to provide the deck, it will be generated for you.")
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
