package main

import (
	"fmt"
)

//GetArrA - Gets reused slice
func GetArrA() []int {
	return []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
}

func main() {
	arrA := GetArrA()
	first()
	fmt.Println(" ")
	second(arrA)
	fmt.Println(" ")
	third(arrA)
	fmt.Println(" ")
	fourth(arrA)
	fmt.Println(" ")
	fifth(arrA)
	fmt.Println(" ")
	sixth()
	fmt.Println(" ")
	seventh()
	mapA := GetMap()
	eigths(mapA)
	fmt.Println(" ")
	nineth(mapA)
	fmt.Println(" ")
	tenth(mapA)
}

func first() {
	arrA := [5]int{1, 2, 3, 4, 5}

	for ind, item := range arrA {
		fmt.Println(ind, item)
	}

	fmt.Printf("%T\n", arrA)
}

func second(arrA []int) {

	for ind, item := range arrA {
		fmt.Println(ind, item)
	}

	fmt.Printf("%T\n", arrA)

}

func third(arrA []int) {
	lastPos := 0
	for ind, item := range arrA {
		if item%5 == 0 {
			fmt.Println(arrA[lastPos : ind+2])
			lastPos = ind + 2
		}
	}

	fmt.Println(arrA[2:7])
	fmt.Println(arrA[1:6])
}

func fourth(x []int) {

	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Println(x)
}

func fifth(x []int) {
	fmt.Println(x)
	x = DeleteFromArray(x, 3, 5)
	fmt.Println(x)
}

// DeleteFromArray - deletes item from slice
//	x - slice of integers to delete from
//	pos - index position in slice (0,1,2,3...)
func DeleteFromArray(x []int, posbeg int, posend int) []int {
	return append(x[:posbeg], x[posend+1:]...)
}

func sixth() {

	x := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
		` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	println("Lenth is", len(x))
	println("Capacity is", cap(x))

	for i := 0; i < 50; i++ {
		println(i, x[i])
	}

}

func seventh() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{a, b}

	for id, item := range x {
		fmt.Println(id, item)
		for idin, itemin := range item {
			fmt.Println("\t", idin, itemin)
		}
	}
}

//GetMap - gets the reused data in map
func GetMap() map[string][]string {
	return map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

}

func eigths(a map[string][]string) {
	for key, value := range a {
		fmt.Println(key, " likes:")
		for idin, itemin := range value {
			fmt.Printf("\t%v\t%v\n", idin, itemin)
		}
	}
}

func nineth(a map[string][]string) {
	a["chainsaws"] = []string{"Pepsi", "Coding", "Chocolates", "Games"}

	eigths(a)
}

func tenth(a map[string][]string) {
	delete(a, "chainsaws")

	eigths(a)
}
