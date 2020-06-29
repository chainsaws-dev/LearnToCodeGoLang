package word

import (
	"fmt"
	"myprojects/013-level-13/02/quote"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "Test 01",
			args: args{s: "one two three three three one"},
			want: map[string]int{
				"one":   2,
				"two":   1,
				"three": 3,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 01",
			args: args{s: "one two three four"},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three four"))
	// Output:
	// 4
}

func BenchmarkCount(t *testing.B) {
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(t *testing.B) {
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		UseCount(quote.SunAlso)
	}
}
