package main

import (
	"fmt"
	"time"
)

var curyear, _, _ = time.Now().Date()

func main() {

	//numx := 256
	//fmt.Println(printdecbinhex(numx))

	// comparisonpractice()

	// constantprint()

	//shiftofbits()

	//rawstringliteral()

	printyears()
}

//Assignment #01
func printdecbinhex(x int) string {
	return fmt.Sprintf("%d\t%b\t%#x\n", x, x, x)
}

//Assignment #02
func comparisonpractice() {
	test1 := 0 == 0
	test2 := 12 <= 0
	test3 := -12 >= 0
	test4 := 2 != 0
	test5 := 3 < 0
	test6 := -3 > 0

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
	fmt.Println(test6)

}

//Assignment #03
const (
	a     = 32
	b int = 42
)

func constantprint() {
	fmt.Println(a, b)
}

//Assignment #04
func shiftofbits() {
	intx := 256
	fmt.Printf("%d\t%b\t%#x\n", intx, intx, intx)

	intmod := intx << 1
	fmt.Printf("%d\t%b\t%#x\n", intmod, intmod, intmod)
}

//Assignment #05
func rawstringliteral() {
	strt := `Это сырой 
	литерал строки
			Как Маяковский
					херачил стихи`

	fmt.Println(strt)
}

//Assignment #06
const (
	y    = iota
	ny   = iota
	nny  = iota
	nnny = iota
)

func printyears() {
	fmt.Println(curyear+y, curyear+ny, curyear+nny, curyear+nnny)
}
