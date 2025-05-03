package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	deck_utils "example.com/go-pontifex/pkg/deck_utils"
)

var suit = [4]string{"clubs", "diamonds", "hearts", "spades"}

var rank = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// A struct for response parsing at /cipher
type CipherResponse struct {
	Answer string    `json:"answer"`
	Deck   *[]string `json:"deck"`
}

// A struct for request parsing at /cipher
type CipherRequest struct {
	Message string   `json:"message"`
	Deck    []string `json:"deck"`
}

// A struct for response parsing at /decipher
type DecipherResponse struct {
	Answer string   `json:"answer"`
	Deck   []string `json:"deck"`
}

// A struct for request parsing at /decipher
type DecipherRequest struct {
	Message string   `json:"message"`
	Deck    []string `json:"deck"`
}

// A struct for request parsing at /generate
type GenerateDeckResponse struct {
	Deck []string `json:"deck"`
}

// A handler for index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	file, err := os.Open("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

// A handler for /cipher page
func CipherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	var inputData CipherRequest

	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(inputData.Message) == 0 {
		http.Error(w, "Запрос не содержит сообщения", http.StatusBadRequest)
		return
	}

	if len(inputData.Deck) == 0 {
		inputData.Deck = deck_utils.DeckShuffle(deck_utils.DeckGenerator(suit, rank))
	}

	initialDeck := make([]string, len(inputData.Deck))
	copy(initialDeck, inputData.Deck)

	var cipherTextAnswer = deck_utils.CipherText(inputData.Message, inputData.Deck)

	// Устанавливаем заголовок ответа как JSON
	w.Header().Set("Content-Type", "application/json")

	// Формируем ответ
	resp := CipherResponse{Answer: cipherTextAnswer, Deck: &initialDeck}

	// Отправляем JSON-ответ
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// A handler for /decipher page
func DecipherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	var inputData DecipherRequest

	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(inputData.Message) == 0 {
		http.Error(w, "Запрос не содержит сообщения", http.StatusBadRequest)
		return
	}

	if len(inputData.Deck) == 0 {
		http.Error(w, "Запрос не содержит колоду", http.StatusBadRequest)
		return
	}

	initialDeck := make([]string, len(inputData.Deck))
	copy(initialDeck, inputData.Deck)

	var cipherTextAnswer = deck_utils.DecipherText(inputData.Message, inputData.Deck)

	// Устанавливаем заголовок ответа как JSON
	w.Header().Set("Content-Type", "application/json")

	// Формируем ответ
	resp := DecipherResponse{Answer: cipherTextAnswer, Deck: initialDeck}

	// Отправляем JSON-ответ
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// A handler for /generate page
func GenerateDeckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}
	var outputData GenerateDeckResponse
	outputData.Deck = deck_utils.DeckShuffle(deck_utils.DeckGenerator(suit, rank))

	w.Header().Set("Content-Type", "application/json")

	resp := GenerateDeckResponse{Deck: outputData.Deck}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
