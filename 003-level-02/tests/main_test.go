package main

import (
	"testing"
)

//Move file to root folder of project to test
func Test_printdecbinhex(t *testing.T) {
	type args struct {
		x int
	}
	type singletest struct {
		name string
		args args
		want string
	}

	tests := []singletest{}

	test01 := singletest{}
	test01.name = "Test 01"
	test01.args.x = 256
	test01.want = "256\t100000000\t0x100\n"

	tests = append(tests, test01)

	test02 := singletest{}
	test02.name = "Test 02"
	test02.args.x = 1024
	test02.want = "1024\t10000000000\t0x400\n"

	tests = append(tests, test02)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printdecbinhex(tt.args.x); got != tt.want {
				t.Errorf("printdecbinhex() = %v,\n want %v", got, tt.want)
			}
		})
	}
}
