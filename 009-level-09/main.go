package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//first()
	//second()
	//third()
	//fourth()
	fifth()
}

var wg sync.WaitGroup

func first() {
	wg.Add(2)
	fmt.Println(runtime.GOOS, runtime.GOARCH, "CPU:", runtime.NumCPU())
	fmt.Println()

	go func() {
		fmt.Println("WARNING!")
		wg.Done()
	}()

	go func() {
		fmt.Println("GOPHERS DISPATCHED TO YOUR LOCATION!")
		wg.Done()
	}()

	fmt.Println("Main finished.")

	wg.Wait()
}

type human interface {
	Speak()
}

type person struct {
	FName      string
	LName      string
	Age        int
	CatchFrase string
}

func (p *person) Speak() {
	fmt.Println(p.FName, p.LName, "says", p.CatchFrase)
}

func second() {

	p := person{
		FName:      "Anonimus",
		LName:      "Mask",
		Age:        32,
		CatchFrase: "Somebody stop me!",
	}

	//cannot use p (type person) as type human in argument to SaySomething:
	// person does not implement human (Speak method has pointer receiver)
	//	SaySomething(p)
	// It is detected by linter even ^_^
	SaySomething(&p)

}

// SaySomething - Вызывает метод Speak
func SaySomething(h human) {
	h.Speak()
}

func third() {
	counter := 0

	fmt.Println("CPUS: ", runtime.NumCPU())

	var wg sync.WaitGroup

	maxgoroutines := 50

	wg.Add(maxgoroutines)

	var mu sync.Mutex

	for i := 0; i < maxgoroutines; i++ {
		go func() {
			mu.Lock()
			v := counter

			//runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter: ", counter)
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("Counter: ", counter)
}

func fourth() {
	var counter int64

	fmt.Println("CPUS: ", runtime.NumCPU())

	var wg sync.WaitGroup

	maxgoroutines := 50

	wg.Add(maxgoroutines)

	for i := 0; i < maxgoroutines; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))

			wg.Done()
		}()

	}

	wg.Wait()

}

func fifth() {
	fmt.Println("OS:", runtime.GOOS, "Architecture:", runtime.GOARCH)
}
