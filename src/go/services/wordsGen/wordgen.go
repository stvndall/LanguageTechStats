package wordsGen

var words = []string{"ambition", "delicious", "national", "penguin", "falling", "cords", "bleeding", "bumble", "devices", "hyaena"}

func FindSequence(count int) []string {
	length := len(words)
	fullCopy := count / length
	left := count % length
	returning := words[:left]
	for i := 0; i <= fullCopy; i++ {
		returning = append(returning, words...)
	}
	return returning
}
