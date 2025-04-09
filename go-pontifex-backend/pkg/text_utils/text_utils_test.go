package text_utils

import (
	"fmt"
	"strings"
	"testing"
)

var alphabet = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

// var numberedText = []string{"1", "2", "3", "4", "5"}

func TestTextToNumber(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyz"
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToUpper(text)
	// numbers := []int{}
	fmt.Println(text)
	if text != "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		t.Error("Chars are wrong!")
	}
	numbers := []int{}
	for i := range text {
		value := alphabet[string(text[i])]
		numbers = append(numbers, value)
	}
	fmt.Println(numbers)
	if len(text) != len(numbers) {
		t.Error("Length of arrays is not equal")
	}
	target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	for i := range numbers {
		if numbers[i] != target[i] {
			t.Error("output numbers are wrong")
		}
	}

}

func TestNumberToKey(t *testing.T) {
	numberedText := []int{1, 26}
	keyStream := []int{1, 53}
	keyes := NumberToKey(numberedText, keyStream)
	fmt.Println("NUMBER TO KEY ==", keyes)
	testCase := []int{2, 1}
	for i := range testCase {
		if keyes[i] != testCase[i] {
			t.Error("input data is out of bounds")
		}
	}
}

func TestKeyToNumber(t *testing.T) {
	numberedText := []int{2, 26}
	keyStream := []int{1, 53}
	keyes := KeyToNumber(numberedText, keyStream)
	testCase := []int{1, 25}
	for i := range testCase {
		if keyes[i] != testCase[i] {
			t.Error("input data is out of bounds")
		}
	}
}

func TestKeyToText(t *testing.T) {}
