package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CipherDecipherResponse struct {
	Message string   `json:"answer"`
	Deck    []string `json:"deck"`
}

func HandleCipherCommand(message string, deck string) []string {
	url := "http://pntfx-backend:8080/cipher"

	deckArr := []string{}
	if deck == "" || deck == "no deck" {
		deckArr = []string{}
	} else {
		deckArr = strings.Split(deck, " ")
	}

	type CipherRequest struct {
		Message string   `json:"message"`
		Deck    []string `json:"deck"`
	}

	data := CipherRequest{message, deckArr}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Ошибка!", err)
		return []string{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []string{}
	}

	var result CipherDecipherResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []string{}
	}

	str := []string{}

	str = append(str, result.Message)
	str = append(str, result.Deck...)

	return str
}

func HandleDecipherCommand(message string, deck string) []string {
	url := "http://pntfx-backend:8080/decipher"

	deckArr := strings.Split(deck, " ")

	type DecipherRequest struct {
		Message string   `json:"message"`
		Deck    []string `json:"deck"`
	}

	data := DecipherRequest{message, deckArr}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка!", err)
		return []string{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return []string{}
	}

	var result CipherDecipherResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return []string{}
	}

	str := []string{}

	str = append(str, result.Message)
	str = append(str, result.Deck...)

	return str

}

func HandleGenerateCommand() string {

	type GenerateDeckResponse struct {
		Deck []string `json:"deck"`
	}

	url := "http://pntfx-backend:8080/generate"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ошибка!", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return ""
	}

	var result GenerateDeckResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return ""
	}
	str := ""

	for _, card := range result.Deck {
		str += card + " "
	}

	return str
}
