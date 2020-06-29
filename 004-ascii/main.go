package main

import "fmt"

func main() {

	fmt.Println("binary\tdecimal\thex\tcharacter")

	for i := 32; i < 127; i++ {
		fmt.Printf("%b\t%d\t%#x\t%c\n", i, i, i, i)
	}

}
