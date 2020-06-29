// Package mymath - contains custom math functions
package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(t *testing.B) {
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		CenteredAvg([]int{10, 1, 4, 2, 6})
	}
}

func TestCenteredAvg(t *testing.T) {
	type args struct {
		xi []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Test 01",
			args: args{xi: []int{10, 1, 4, 2, 6}},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CenteredAvg(tt.args.xi); got != tt.want {
				t.Errorf("CenteredAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{100, 1, 4, 2, 6}))
	// Output:
	// 4
}
