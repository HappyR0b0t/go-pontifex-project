package deck_utils

import (
	"reflect"
	"testing"

	"example.com/go-pontifex/pkg/utils"
)

// var a = [4]string{"clubs", "diamonds", "hearts", "spades"}
// var b = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// func TestDeckGenerator(t *testing.T) {

// 	got := DeckGenerator(a, b)
// 	want := 1
// 	if got["clubs-A"] != want {
// 		t.Error("Key and value are incorrect")
// 	}
// 	if got["JA"] != 53 {
// 		t.Error("Key and value are incorrect")
// 	}
// 	if got["JB"] != 53 {
// 		t.Error("Key and value are incorrect")
// 	}
// 	if len(got) != 54 {
// 		t.Error("Array length is incorrect")
// 	}
// }

func TestDeckShuffle(t *testing.T) {
	// got := DeckGenerator(a, b)
	// want := DeckShuffle(got)
	// array := []string{}

	// array = append(array, want...)
	// if len(got) != len(array) {
	// 	t.Error("Length of decks is not equal")
	// }
	// i := 0
	// for k := range got {
	// 	fmt.Println(k, array[i])
	// 	if k != array[i] {
	// 		t.Error("Decks are identical")
	// 	}
	// 	i++
	// }
}

func TestMoveJocker(t *testing.T) {
	tests := []struct {
		name              string
		inputDeck         []string
		inputCurrentIndex int
		inputTargetIndex  int
		expected          []string
	}{
		{
			name:              "Puts last card on top of the deck",
			inputDeck:         []string{"JA", "a", "b", "JB"},
			inputCurrentIndex: 3,
			inputTargetIndex:  0,
			expected:          []string{"JB", "JA", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MoveJocker(tt.inputDeck, tt.inputCurrentIndex, tt.inputTargetIndex)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MoveJocker() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindJocker(t *testing.T) {
	tests := []struct {
		name        string
		inputDeck   []string
		inputString string
		expected    int
	}{
		{
			name:        "Find jocker A",
			inputDeck:   []string{"a", "b", "c", "JA", "d", "e", "JB"},
			inputString: "JA",
			expected:    3,
		},
		{
			name:        "Find jocker B",
			inputDeck:   []string{"a", "b", "c", "JA", "d", "e", "JB"},
			inputString: "JB",
			expected:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindJocker(tt.inputDeck, tt.inputString)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindJocker() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMoveJockerA(t *testing.T) {
	tests := []struct {
		name          string
		inputDeck     []string
		inputIndex    int
		expectedDeck  []string
		expectedIndex int
	}{
		{
			name:          "Move jocker A, jocker A is the last card",
			inputDeck:     []string{"a", "b", "c", "JB", "d", "e", "JA"},
			inputIndex:    6,
			expectedDeck:  []string{"a", "JA", "b", "c", "d", "e", "JB"},
			expectedIndex: 1,
		},
		{
			name:          "Move jocker A, jocker A is NOT the last card",
			inputDeck:     []string{"a", "b", "c", "JA", "d", "e", "JB"},
			inputIndex:    3,
			expectedDeck:  []string{"a", "b", "c", "d", "JA", "e", "JB"},
			expectedIndex: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, resultIndex := MoveJockerA(tt.inputDeck, tt.inputIndex)
			if !reflect.DeepEqual(result, tt.expectedDeck) && !reflect.DeepEqual(resultIndex, tt.expectedIndex) {
				t.Errorf("FindJocker() = %v, want %v", result, tt.expectedDeck)
			}
		})
	}
}

func TestMoveJockerB(t *testing.T) {
	tests := []struct {
		name          string
		inputDeck     []string
		inputIndex    int
		expectedDeck  []string
		expectedIndex int
	}{
		{
			name:          "Move jocker B, jocker B is the last card",
			inputDeck:     []string{"a", "b", "c", "JA", "d", "e", "JB"},
			inputIndex:    6,
			expectedDeck:  []string{"a", "b", "JB", "c", "JA", "d", "e"},
			expectedIndex: 2,
		},
		{
			name:          "Move jocker B, jocker B is the second last card",
			inputDeck:     []string{"a", "b", "c", "d", "e", "JB", "JA"},
			inputIndex:    5,
			expectedDeck:  []string{"a", "JB", "b", "c", "d", "e", "JA"},
			expectedIndex: 1,
		},
		{
			name:          "Move jocker B, jocker B is NOT the last card",
			inputDeck:     []string{"a", "b", "c", "JB", "JA", "d", "e"},
			inputIndex:    3,
			expectedDeck:  []string{"a", "b", "c", "JA", "d", "JB", "e"},
			expectedIndex: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultDeck, resultIndex := MoveJockerB(tt.inputDeck, tt.inputIndex)
			if !reflect.DeepEqual(resultDeck, tt.expectedDeck) || !reflect.DeepEqual(resultIndex, tt.expectedIndex) {
				t.Errorf("FindJocker() = %v, result index = %v, want %v, wanted index is %v", resultDeck, resultIndex, tt.expectedDeck, tt.expectedIndex)
			}
		})
	}
}

func TestJockerShift(t *testing.T) {

}

func TestTripleCut(t *testing.T) {
	tests := []struct {
		name         string
		inputDeck    []string
		inputIndices []int
		expected     []string
	}{
		{
			name:         "Top of the deck is empty",
			inputDeck:    []string{"JA", "a", "JB", "b"},
			inputIndices: []int{0, 2},
			expected:     []string{"b", "JA", "a", "JB"},
		},
		{
			name:         "Bottom of the deck is empty",
			inputDeck:    []string{"a", "JA", "b", "JB"},
			inputIndices: []int{1, 3},
			expected:     []string{"JA", "b", "JB", "a"},
		},
		{
			name:         "Jockers are adjacent, top of the deck is empty ",
			inputDeck:    []string{"JA", "JB", "a", "b"},
			inputIndices: []int{0, 1},
			expected:     []string{"a", "b", "JA", "JB"},
		},
		{
			name:         "Jockers are adjacent, bottom of the deck is empty ",
			inputDeck:    []string{"a", "b", "JA", "JB"},
			inputIndices: []int{2, 3},
			expected:     []string{"JA", "JB", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TripleCut(tt.inputDeck, tt.inputIndices)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TripleCut() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountCut(t *testing.T) {
	tests := []struct {
		name       string
		inputDeck  []string
		inputValue int
		expected   []string
	}{
		{
			name:       "Count cut: cut after first card, value equals 1",
			inputDeck:  []string{"a", "b", "c", "d", "e", "f"},
			inputValue: 1,
			expected:   []string{"b", "c", "d", "e", "a", "f"},
		},
		{
			name:       "Count cut: value equals second last card, value equals 53",
			inputDeck:  []string{"a", "b", "c", "d", "e", "f"},
			inputValue: 5,
			expected:   []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountCut(tt.inputDeck, tt.inputValue)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountCut() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCardToNumber(t *testing.T) {

}

func TestFindOutput(t *testing.T) {}

func TestKeyStream(t *testing.T) {
	numberedText := []int{1, 2, 3}
	textLength := len(numberedText)
	inputDeck := utils.ReadDeck("test_input_deck_one.txt")
	_, keyStream := KeyStream(textLength, &inputDeck)
	if len(keyStream) == 0 {
		t.Error("Keystream array length is zero!")
	}
}

func TestKeyStreamRecusrsive(t *testing.T) {}
