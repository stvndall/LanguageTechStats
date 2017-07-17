package letterCounts

import "sync"

func countLetters(letter rune, letterCounts map[rune]int) {
	letterCounts[letter] = letterCounts[letter] + 1
}

func forEachLetter(letters chan rune, counts map[rune]int, finishedSignal * sync.WaitGroup) {
	for l := range letters {
		countLetters(l, counts)
	}
	finishedSignal.Done()
}


func countLettersInWord(words chan string) map[rune]int {
	runeCheck := make(chan rune)
	letterCount := make(map[rune]int)
	var completeSignal sync.WaitGroup
	completeSignal.Add(1)
	go forEachLetter(runeCheck, letterCount, &completeSignal)
	for word := range words {
		for _, l := range word {
			runeCheck <- l
		}
	}
	close(runeCheck)

	completeSignal.Wait()
	return letterCount
}
