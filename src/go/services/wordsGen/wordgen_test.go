package wordsGen

import "testing"

func TestFindSequenceOf10(t *testing.T) {
	RunTest := func(count int) {
		arr := FindSequence(count)
		length := len(arr)
		if length != count {
			t.Errorf("incorrect returned Amount, expected %d, received %d", count, length)
		}
	}
	RunTest(0)
	RunTest(9)
	RunTest(19)
	RunTest(100)
	RunTest(105)
}
