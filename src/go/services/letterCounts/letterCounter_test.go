package letterCounts

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountLettersWhenCountAlreadyExists(t *testing.T) {
	expected := 5
	letters := make(map[rune]int)
	letters['A'] = 4
	countLetters('A', letters)
	assert.Equal(t, letters['A'], expected)
}

func TestCountLetters_WhenCountDoesNotExist(t *testing.T) {
	expectedB := 2
	expectedA := 4
	letters := make(map[rune]int)
	letters['A'] = 4
	countLetters('B', letters)
	countLetters('B', letters)
	assert.Equal(t,  letters['A'], expectedA)
	assert.Equal(t,  letters['B'], expectedB)
	
}

func TestCountLetterInSingleWord(t *testing.T) {
	words := make(chan string, 1)
	words <- "aabc"
	close(words)
	letterCount := countLettersInWord(words)
	assert.Equal(t, letterCount['a'], 2, "" )
	assert.Equal(t, letterCount['b'], 1, "" )
	assert.Equal(t, letterCount['c'], 1, "expected run c to have a count of 1" )
}
