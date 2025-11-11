package deck_utils

import (
	"math/rand"
	"strings"

	"example.com/go-pontifex/pkg/text_utils"

	"github.com/rs/zerolog/log"
)

var suitsMap = map[string]int{
	"clubs":    0,
	"diamonds": 13,
	"hearts":   26,
	"spades":   39,
}

var rankMap = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"JA": 53,
	"JB": 53,
}

func DeckGenerator(suit [4]string, rank [13]string) []string {
	deck := []string{}
	for _, i := range suit {
		for _, j := range rank {
			deck = append(deck, i+"-"+j)
		}
	}
	deck = append(deck, "JA")
	deck = append(deck, "JB")
	return deck
}

// Shuffles the deck `randomly`
func DeckShuffle(deck []string) []string {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

// Moves Jocker to target position
func MoveJocker(deckKeys []string, current int, target int) []string {
	for i := current; i > target; i-- {
		deckKeys[i-1], deckKeys[i] = deckKeys[i], deckKeys[i-1]
	}
	return deckKeys
}

// Finds a jocker in a deck
func FindJocker(deck []string, jocker string) int {
	for i := range deck {
		if deck[i] == jocker {
			return i
		}
	}
	return -1
}

// Move Jocker A
func MoveJockerA(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == len(deck)-1 {
		MoveJocker(deck, i, 1)
		jockerIndex = 1
	} else {
		MoveJocker(deck, i, i+1)
		jockerIndex = i + 1
	}
	return deck, jockerIndex
}

// Move Jocker B
func MoveJockerB(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == len(deck)-1 {
		MoveJocker(deck, i, 2)
		jockerIndex = 2
	} else if i == len(deck)-2 {
		MoveJocker(deck, i, 1)
		jockerIndex = 1
	} else {
		deck[i], deck[i+1], deck[i+2] = deck[i+1], deck[i+2], deck[i]
		jockerIndex = i + 2
	}
	return deck, jockerIndex
}

// Shifts both jokers accordingly
func JockerShift(deckKeys []string) ([]string, []int) {
	jockers := []int{}
	ja := FindJocker(deckKeys, "JA")
	deckKeys, _ = MoveJockerA(deckKeys, ja)
	jb := FindJocker(deckKeys, "JB")
	deckKeys, jb = MoveJockerB(deckKeys, jb)
	ja = FindJocker(deckKeys, "JA")
	jockers = append(jockers, ja, jb)
	if jockers[0] > jockers[1] {
		jockers[0], jockers[1] = jockers[1], jockers[0]
	}
	return deckKeys, jockers
}

// Performs a triple cut on a deck
func TripleCut(deckKeys []string, jockers []int) []string {
	deck := []string{}
	top := deckKeys[:jockers[0]]
	middle := deckKeys[jockers[0] : jockers[1]+1]
	bottom := deckKeys[jockers[1]+1:]

	deck = append(deck, bottom...)
	deck = append(deck, middle...)
	deck = append(deck, top...)

	return deck
}

// Performs a count cut on a deck
func CountCut(tripleCutDeck []string, value int) []string {
	lastIndex := len(tripleCutDeck) - 1

	top := tripleCutDeck[:value]
	middle := tripleCutDeck[value:lastIndex]
	bottom := tripleCutDeck[lastIndex]

	countCutDeck := []string{}
	countCutDeck = append(countCutDeck, middle...)
	countCutDeck = append(countCutDeck, top...)
	countCutDeck = append(countCutDeck, bottom)

	return countCutDeck
}

// Converts a card to number
func cardToNumber(card string) int {
	if card == "JA" || card == "JB" {
		return 53
	} else {
		suitAndRank := strings.Split(card, "-")
		return suitsMap[suitAndRank[0]] + rankMap[suitAndRank[1]]
	}
}

// Finds output card in a deck
func FindOutput(tripleCutDeck []string) int {
	return cardToNumber(tripleCutDeck[0])

}

// Creates a keystream for conversion into chars non-recursively
func KeyStream(textLength int, inputDeck *[]string) ([]string, []int) {
	var keyStream = &[]int{}
	for i := 0; i < textLength; {
		jockers := &[]int{}
		lastIndex := len(*inputDeck) - 1

		*inputDeck, *jockers = JockerShift(*inputDeck)
		*inputDeck = TripleCut(*inputDeck, *jockers)
		lastCardValue := cardToNumber((*inputDeck)[lastIndex])
		*inputDeck = CountCut(*inputDeck, lastCardValue)
		key := FindOutput(*inputDeck)
		if key == 53 {
			log.Info().
				Str("function", "KeyStream").
				Int("key", key).
				Msg("key was a joker")
			continue
		}
		*keyStream = append(*keyStream, key)
		i++
	}

	return *inputDeck, *keyStream
}

// A function to cipher provided text with provided deck
func CipherText(plainText string, inputDeck []string) string {
	numberedText := text_utils.TextToNumber(plainText)
	var textLength int = len(numberedText)
	_, keyStream := KeyStream(textLength, &inputDeck)
	keys := text_utils.NumberToKey(numberedText, keyStream)
	cipheredText := text_utils.KeyToText(keys)

	return cipheredText
}

// A function to decipher provided text with provided deck
func DecipherText(cipheredText string, inputDeck []string) string {
	numberedText := text_utils.TextToNumber(cipheredText)
	var textLength int = len(numberedText)
	_, keyStream := KeyStream(textLength, &inputDeck)
	keys := text_utils.KeyToNumber(numberedText, keyStream)
	decipheredText := text_utils.KeyToText(keys)

	return decipheredText
}
