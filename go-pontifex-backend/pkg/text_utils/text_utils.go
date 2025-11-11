package text_utils

import (
	"strings"

	"github.com/rs/zerolog/log"
)

// Converts text to numbers
func TextToNumber(text string) []int {
	text = strings.ReplaceAll(strings.ToUpper(text), " ", "")
	numbers := []int{}
	var c int = 64
	for _, val := range []byte(text) {
		numbers = append(numbers, int(val)-c)
	}
	return numbers
}

// Converts numbers to key for the keystream
func NumberToKey(numberedText []int, keyStream []int) []int {
	if len(numberedText) != len(keyStream) {
		log.Error().
			Str("function", "NumberToKey").
			Int("lenNumberedText", len(numberedText)).
			Int("lenKeyStream", len(keyStream)).
			Msg("Mismatched argument lengths")
		return nil
	}

	keys := []int{}
	m := 26
	for i := range numberedText {
		n := numberedText[i] + keyStream[i]
		if n%m == 0 {
			keys = append(keys, 26)
		} else {
			keys = append(keys, n%m)
		}
	}
	return keys
}

// Converts keys from the keystream to numbers
func KeyToNumber(numberedText []int, keyStream []int) []int {
	if len(numberedText) != len(keyStream) {
		log.Error().
			Str("function", "KeyToNumber").
			Int("lenNumberedText", len(numberedText)).
			Int("lenKeyStream", len(keyStream)).
			Msg("Mismatched argument lengths")
		return nil
	}
	keys := []int{}
	for i := range numberedText {
		m := 26
		if numberedText[i] < keyStream[i]%m {
			n := (numberedText[i] + m) - keyStream[i]%m
			keys = append(keys, n)
		} else {
			n := (numberedText[i]) - keyStream[i]%m
			keys = append(keys, n)
		}
	}
	return keys
}

// V2 Converts keys from the keystream to chars
func KeyToText(keys []int) string {
	var text string
	var c int = 64
	for i, val := range keys {
		if i%5 == 0 && i > 4 {
			text += " "
		}
		text += string(byte(val + c))
	}
	return text
}
