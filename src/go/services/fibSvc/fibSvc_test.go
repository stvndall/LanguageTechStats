package fibSvc

import "testing"

func TestFibFindFirst(t *testing.T) {
	first, second := generateFirst()
	if first != 0 || second != 1 {
		t.Error("unexpected first records")
	}
}

func TestGenerateNext(t *testing.T) {
	expected := 3
	next := generateNext(1, 2)

	if next != expected {
		t.Errorf("expected %d, but got %d", expected, next)
	}
}

func TestGetHighestBefore(t *testing.T) {
	testReturned := func(t *testing.T, number, expectedIndex, expectedValue int) {
		index, indexValue := getHighestBefore(number)
		if index != expectedIndex {
			t.Errorf("Expected index %d, but got %d", expectedIndex, index)
		}
		if expectedValue != indexValue {
			t.Errorf("Expected Value %d, but got %d", expectedValue, indexValue)
		}
	}

	t.Run("low and appears on fib sequence", func(t *testing.T) { testReturned(t, 8, 7, 8) })
	t.Run("low and does not appears on fib sequence", func(t *testing.T) { testReturned(t, 6, 6, 5) })
	t.Run("high and appears on fib sequence", func(t *testing.T) { testReturned(t, 6765, 21, 6765) })
	t.Run("high and does not appears on fib sequence", func(t *testing.T) { testReturned(t, 14112, 22, 10946) })

}
