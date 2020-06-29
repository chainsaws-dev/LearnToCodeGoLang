package main

import (
	"fmt"
)

type gogo int

var xxx gogo
var yyy int

func fourth() {
	fmt.Println("\n\tExercize 04")

	fmt.Println(xxx)
	fmt.Printf("%T\n", xxx)
	xxx = 42
	fmt.Println(xxx)

	yyy = int(xxx)
	fmt.Println(yyy)
	fmt.Printf("%T\n", yyy)
}
