package average

import "testing"

func Test_averageNumbers(t *testing.T) {
	tests := []struct {
		name        string
		allNumbers        []int
		wantAverage float64
		wantErr     bool
	}{
	{"given no numbers should error", []int{}, 0, true},
	{"given 1 number should return same number", []int{124}, 124, false},
	{"given multiple numbers should return average", []int{2,2,4,0}, 2, false},
	{"given multiple numbers with unreal average should return average", []int{557,112,1248,1241,1}, 631.8, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAverage, err := averageNumbers(tt.allNumbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("averageNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAverage != tt.wantAverage {
				t.Errorf("averageNumbers() = %v, want %v", gotAverage, tt.wantAverage)
			}
		})
	}
}
