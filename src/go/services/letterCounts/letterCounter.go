package letterCounts

import "sync"
import "fmt"

func countLetters(letter rune, letterCounts map[rune]int) {
	letterCounts[letter] = letterCounts[letter] + 1
}

func forEachLetter(letters chan rune, counts map[rune]int, finishedSignal *sync.WaitGroup) {
	for l := range letters {
		countLetters(l, counts)
	}
	finishedSignal.Done()
}
func breakWords(word string, letterChan chan rune) {
	for _, l := range word {
		letterChan <- l
	}
}

func countLettersInWord(words chan string) map[rune]int {
	fmt.Printf("here")
	runeCheck := make(chan rune)
	letterCount := make(map[rune]int)
	var completeSignal sync.WaitGroup
	completeSignal.Wait(1)
	go forEachLetter(runeCheck, letterCount, &completeSignal)
	fmt.Printf("letters started")
	for word := range words {
		breakWords(word, runeCheck)
	}
	fmt.Printf("letters complete")
	close(runeCheck)

	completeSignal.Wait()
	return letterCount
}
