package main

import (
	"fmt"
)

func main() {
	//first()
	second()
}

func first() {
	a := "Precios code: 1111 2222 3333 4444"

	fmt.Println(&a)
}

type person struct {
	firstname  string
	secondname string
	familyname string
	age        int
}

func second() {
	a := person{
		firstname:  "Иванов",
		secondname: "Иван",
		familyname: "Иванович",
		age:        44,
	}

	fmt.Println(a)

	changeme(&a)

	fmt.Println(a)

}

func changeme(perstochange *person) {
	(*perstochange).age = 22
}
