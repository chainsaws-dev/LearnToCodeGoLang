// Implements solutions for level 11 tasks
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	//first()
	//second()
	//third()
	//fourth()
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

/*
//!+first
	first()
//!-first
*/

func first() {
	//!+first
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
	//!-first
}

func second() {
	//!+second
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
	//!-first
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("Error occured %v when tried to marshal %v", err.Error(), a)
	}

	return bs, nil
}

type customErr struct {
	ErrorCode string
	ErrorText string
	Critical  bool
}

func (err customErr) Error() string {
	var errstring string

	if err.Critical {
		if len(err.ErrorCode) == 0 {
			errstring = fmt.Sprintf("Critical error occured: %v", err.ErrorText)
		} else {
			errstring = fmt.Sprintf("Critical error № %v occured: %v", err.ErrorCode, err.ErrorText)
		}
	} else {
		if len(err.ErrorCode) == 0 {
			errstring = fmt.Sprintf("Error occured: %v", err.ErrorText)
		} else {
			errstring = fmt.Sprintf("Error № %v occured: %v", err.ErrorCode, err.ErrorText)
		}
	}

	return errstring
}

func third() {
	//!+third
	c1 := customErr{
		ErrorCode: "101",
		ErrorText: "I am a teapot!",
		Critical:  true,
	}
	foo(c1)
	//!+third
}

func foo(e error) {
	fmt.Println(e)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func fourth() {
	//!+fourth
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)

	_, errsqrt := sqrt(-10.23)
	if errsqrt != nil {
		log.Println(errsqrt)
	}
	//!-fourth
}

func sqrt(f float64) (float64, error) {
	if f < 0 {

		se := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("Incorrect value %v", f),
		}

		return -1, se
	}
	return 42, nil
}

func fifth() {

}
