package letterCounts

import "sync"
import "fmt"

func countLetters(letter rune, letterCounts map[rune]int) {
	letterCounts[letter] = letterCounts[letter] + 1
}

func forEachLetter(letters chan rune, counts map[rune]int, finishedSignal sync.WaitGroup) {
	for l := range letters {
		countLetters(l, counts)
	}
	finishedSignal.Done()
}
func breakWords(word string, counts map[rune]int) {
	for _, l := range word {
		countLetters(l, counts)
	}
}

func countLettersInWord(words chan string) map[rune]int {
	fmt.Printf("here")
	// runeCheck := make(chan rune)
	letterCount := make(map[rune]int)
	// var completeSignal sync.WaitGroup
	// completeSignal.Add(1)
	// go forEachLetter(runeCheck, letterCount, completeSignal)
	fmt.Printf("letters started")
	for word := range words {
		breakWords(word, letterCount)
	}
	fmt.Printf("letters complete")
	// close(runeCheck)

	// completeSignal.Wait()
	return letterCount
}
