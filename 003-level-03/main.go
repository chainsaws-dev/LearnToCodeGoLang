package main

import "fmt"

func main() {
	//first()
	//second()
	//third()
	//forth()
	//fifth()
	//sixthseventh(2, 3)
	//sixthseventh(4, 4)
	//sixthseventh(3, 1)
	fmt.Println(eigth("Running"))
	fmt.Println(eigth("Skiing"))
	fmt.Println(eigth("John"))

}

func first() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func second() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U\t%c\n", i, i)
		}
	}
}

func third() {
	bd := 1989
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}

func forth() {
	bd := 1989
	for {
		if bd > 2020 {
			break
		}

		fmt.Println(bd)
		bd++
	}
}

func fifth() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func sixthseventh(a int, b int) {
	if a == b {
		fmt.Printf("%v equals %v!\n", a, b)
	} else if a > b {
		fmt.Printf("%v is greater than %v!\n", a, b)
	} else {
		fmt.Printf("%v is less than %v!\n", a, b)
	}
}

func eigth(favSport string) string {
	switch favSport {
	case "Running":
		return "Run, Forrest run!"
	case "Skiing":
		return "Slide, Forrest, slide"
	default:
		return "Not into sports?"
	}
}
