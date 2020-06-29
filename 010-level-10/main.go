package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//firstone()
	//firsttwo()
	//secondone()
	//secondtwo()
	//third()
	//fourth()
	//fifth()
	//sixth()
	seventh()
}

func firstone() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func firsttwo() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func secondone() {
	cs := make(chan int)

	go func(ch chan<- int) {
		ch <- 42
	}(cs)
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func secondtwo() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func third() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(channel <-chan int) {
	for v := range channel {
		fmt.Println(v)
	}
}

func fourth() {
	q := make(chan int)
	c := gen2(q)

	receivefourth(c, q)

	fmt.Println("about to exit")
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receivefourth(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func fifth() {
	c := make(chan int)

	go GetFibonacciRow(50, c)

	v, ok := <-c
	fmt.Println(v, ok)

	for ok {
		v, ok = <-c
		fmt.Println(v, ok)
	}

}

// GetFibonacciRow - Получает ряд Фибоначии до n-ного знака
func GetFibonacciRow(n int, c chan<- int) {

	if n >= 3 {
		var a int
		var b int

		for i := 3; i <= n; i++ {
			if i == 3 {
				a = 1
				b = 1

				c <- 0
				c <- a
				c <- b
			}

			a, b = b, a+b
			c <- b
		}
	} else {
		if n >= 0 {
			for i := 0; i <= n; i++ {
				if i == 0 {
					c <- 0
				} else {
					c <- 1
				}
			}
		} else {
			c <- -1
		}
	}

	close(c)
}

func sixth() {
	c := make(chan int)

	go PutHundred(c)
	GetHundred(c)
}

// PutHundred - сохраняет в канал значения от 1 до 100
func PutHundred(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

// GetHundred - получает значения из канала
func GetHundred(c <-chan int) {

	v, ok := <-c

	for ok {
		if ok {
			fmt.Println(v)
			v, ok = <-c
		}
	}

}

func seventh() {
	nums := make(chan int)

	go RunTen(nums)

	v, ok := <-nums

	for ok {
		if ok {
			fmt.Println(v)
			v, ok = <-nums
		}
	}

}

// RunTen - запускает десять процессов которые добавляют десять чисел
func RunTen(nums chan<- int) {

	for i := 0; i < 10; i++ {
		done := make(chan int)

		go AddTenRandom(i, done)

		for v := range done {
			nums <- v
		}
	}

	close(nums)

}

// AddTenRandom - добавляет десять случайных чисел в канал
func AddTenRandom(proc int, done chan<- int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		done <- r.Int()
	}

	close(done)
}
