package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var standardReqBody = `{
	"message": "Test Message",
	"deck": [
		"clubs-A",
		"clubs-2",
		"clubs-3",
		"clubs-4",
		"clubs-5",
		"clubs-6",
		"clubs-7",
		"clubs-8",
		"clubs-9",
		"clubs-10",
		"clubs-J",
		"clubs-Q",
		"clubs-K",
		"diamonds-A",
		"diamonds-2",
		"diamonds-3",
		"diamonds-4",
		"diamonds-5",
		"diamonds-6",
		"diamonds-7",
		"diamonds-8",
		"diamonds-9",
		"diamonds-10",
		"diamonds-J",
		"diamonds-Q",
		"diamonds-K",
		"hearts-A",
		"hearts-2",
		"hearts-3",
		"hearts-4",
		"hearts-5",
		"hearts-6",
		"hearts-7",
		"hearts-8",
		"hearts-9",
		"hearts-10",
		"hearts-J",
		"hearts-Q",
		"hearts-K",
		"spades-A",
		"spades-2",
		"spades-3",
		"spades-4",
		"spades-5",
		"spades-6",
		"spades-7",
		"spades-8",
		"spades-9",
		"spades-10",
		"spades-J",
		"spades-Q",
		"spades-K",
		"JA",
		"JB"
	]
}`

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	IndexHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Ожидался статус 200, получен %v", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	if len(body) == 0 {
		t.Error("Тело ответа пусто")
	}
}

func TestCipherHandlerNoDeckNoMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/cipher", nil)
	w := httptest.NewRecorder()

	CipherHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != 400 {
		t.Fatalf("Ожидался статус 200, получен %v", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	if len(body) == 0 {
		t.Error("Тело ответа пусто")
	}
}

func TestCipherHandlerWithDeck(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/cipher", strings.NewReader(standardReqBody))
	w := httptest.NewRecorder()

	CipherHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Ожидался статус 200, получен %v", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	if len(body) == 0 {
		t.Error("Тело ответа пусто")
	}
}

func TestCipherHandlerMethodGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cipher", nil)
	w := httptest.NewRecorder()

	CipherHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Ожидался статус 400, получен %v", resp.StatusCode)
	}
}

func TestDecipherHandlerMethodGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/decipher", nil)
	w := httptest.NewRecorder()

	DecipherHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Ожидался статус 400, получен %v", resp.StatusCode)
	}
}

func TestGenerateDeckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/generate", nil)
	w := httptest.NewRecorder()

	DecipherHandler(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, получен %v", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	if len(body) == 0 {
		t.Errorf("Тело ответа пусто: %v", body)
	}
}
