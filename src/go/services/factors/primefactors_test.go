package factors

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantPfs []int
	}{
		{"2",args{2}, []int{2}},
		{"3", args{3}, []int{3}},
		{"97", args{97}, []int{97}},
		{"11212001", args{11212001}, []int{11212001}},
		{"112012", args{112012}, []int{2,2,41,683}},
		{"1121411", args{1121411}, []int{23,48757}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPfs := PrimeFactors(tt.args.number); !reflect.DeepEqual(gotPfs, tt.wantPfs) {
				t.Errorf("PrimeFactors() = %v, want %v", gotPfs, tt.wantPfs)
			}
		})
	}
}
