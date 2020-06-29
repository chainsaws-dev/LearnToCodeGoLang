// Copyright 2020 Alexander Vorobiev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//!+

// Package dog - implements transfer from human years to dog years
package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type args struct {
		HumanYears int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive test",
			args: args{HumanYears: 30},
			want: 210,
		},
		{
			name: "Positive test",
			args: args{HumanYears: 3},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Years(tt.args.HumanYears); got != tt.want {
				t.Errorf("Years() = %v, want %v", got, tt.want)
			}
		})
	}
}

//!+example

func ExampleYears() {
	fmt.Println(Years(30))
	fmt.Println(Years(3))
	// Output:
	// 210
	// 21
}

//!-example
