package main

import (
	"fmt"
	"math"
)

func main() {
	//first()
	//second()
	//fourth()
	//fifth()
	//sixth()
	//seventh()
	//eights()
	nineth()
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 256, "hex"
}

func first() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo2(x ...int) int {
	total := 0
	for _, item := range x {
		total += item
	}

	return total
}

func bar2(x []int) int {
	return foo2(x...)
}

func second() {
	//third assignment
	defer fmt.Println("I am defined on top but will be shown last")

	xx := []int{22, 33, 44, 55, 66, 77, 88, 99}

	xxt := foo2(xx...)

	fmt.Println(xxt)

	xxt = bar2(xx)

	fmt.Println(xxt)
}

type person struct {
	first string
	last  string
	age   int
}

func (per person) speak() {
	fmt.Println(per.first, per.last, per.age)
}

func fourth() {
	x := person{"John", "Lennon", 22}
	x.speak()
}

//fifth
type square struct {
	side int
}

type circle struct {
	radius int
}

func (s square) area() float64 {
	return float64(s.side ^ 2)
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius^2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(int(s.area()))
}

func fifth() {
	sq := square{32}

	cr := circle{16}

	info(sq)
	info(cr)

}

func sixth() {

	func(x string) {
		fmt.Println(x)
	}("It works anonimus!")

}

func seventh() {

	a := func(x string) {
		fmt.Println(x)
	}

	a("It works!")

}

func eights() {
	a := retFuncSmaple()
	a()
}

func retFuncSmaple() func() {
	return func() {
		fmt.Println("EmbeddedFunc")
	}
}

func nineth() {
	a := func(x string) {
		fmt.Println(x)
	}

	argumentfunc(a)
}

func argumentfunc(a func(x string)) {
	a(fmt.Sprintf("binary\tdecimal\thex\tcharacter"))

	for i := 32; i < 127; i++ {
		a(fmt.Sprintf("%b\t%d\t%#x\t%c\n", i, i, i, i))
	}
}
