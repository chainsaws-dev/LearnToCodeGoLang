// Package dog allows us to more fully understand dogs.
package dog

import (
	"fmt"
	"testing"
)

func TestYearsTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{n: 10},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearsTwo(tt.args.n); got != tt.want {
				t.Errorf("Test \"%v\" of function YearsTwo() failed. We got %v and we want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestYears(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{n: 10},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Years(tt.args.n); got != tt.want {
				t.Errorf("Test \"%v\" of function Years() failed. We got %v and we want %v", tt.name, got, tt.want)
			}
		})
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(11))
	// Output:
	// 77
}

func ExampleYears() {
	fmt.Println(Years(11))
	// Output:
	// 77
}

func BenchmarkYearsTwo(t *testing.B) {
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		YearsTwo(11)
	}
}

func BenchmarkYears(t *testing.B) {
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		Years(11)
	}
}
