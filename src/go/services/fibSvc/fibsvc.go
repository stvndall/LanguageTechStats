package fibSvc

func generateFirst() (firstIndex, secondsIndex int) {
	firstIndex = 0
	secondsIndex = 1
	return
}

func generateNext(n1, n2 int) (next int) {
	next = n1 + n2
	return
}

func getHighestBefore(number int) (index, indexValue int) {
	indexValue, upper := generateFirst()
	index = 1
	for upper <= number {
		next := generateNext(indexValue, upper)
		index++
		indexValue = upper
		upper = next
	}
	return
}
