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
	testCases := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "extreme values",
			inputA:   []int{1, 26},
			inputB:   []int{1, 53},
			expected: []int{2, 1},
		},
		{
			name:     "slices of different lengths",
			inputA:   []int{1, 26},
			inputB:   []int{1, 17, 53},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NumberToKey(tc.inputA, tc.inputB)
			for i, v := range got {
				if v != tc.expected[i] {
					t.Errorf("NumberToKey(%d, %d) = %d; want %d", tc.inputA, tc.inputB, got, tc.expected)
				}
			}
		})
	}
}

func TestKeyToNumber(t *testing.T) {
	testCases := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "extreme values",
			inputA:   []int{2, 26},
			inputB:   []int{1, 53},
			expected: []int{1, 25},
		},
		{
			name:     "slices of different lengths",
			inputA:   []int{2, 26},
			inputB:   []int{1, 17, 53},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := KeyToNumber(tc.inputA, tc.inputB)
			for i, v := range got {
				if v != tc.expected[i] {
					t.Errorf("KeyToNumber(%d, %d) = %d; want %d", tc.inputA, tc.inputB, got, tc.expected)
				}
			}
		})
	}
}

func TestKeyToText(t *testing.T) {}
