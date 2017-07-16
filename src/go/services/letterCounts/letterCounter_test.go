package letterCounts

import (
	"testing"
)

func TestCountLettersWhenCountAlreadyExists(t *testing.T) {
	expected := 5
	letters := make(map[rune]int)
	letters['A'] = 4
	countLetters('A', letters)
	if letters['A'] != expected {
		t.Errorf("expected count to be %d but got %d instead", expected, letters['A'])
	}
}

func TestCountLetters_WhenCountDoesNotExist(t *testing.T) {
	expectedB := 1
	expectedA := 4
	letters := make(map[rune]int)
	letters['A'] = 4
	countLetters('B', letters)
	if letters['A'] != expectedA {
		t.Errorf("expected for A count to be %d but got %d instead", expectedA, letters['A'])
	}
	if letters['B'] != expectedB {
		t.Errorf("expected for B count to be %d but got %d instead", expectedB, letters['B'])
	}
}

func TestCountLetterInSingleWord(t *testing.T) {
	t.Errorf("starting")

	words := make(chan string, 1)
	words <- "aabc"
	close(words)
	letterCount := countLettersInWord(words)
	if letterCount['a'] != 2 {
		t.Errorf("A was expected 2 but got %d", letterCount['a'])
	}

}
