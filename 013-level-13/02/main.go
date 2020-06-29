// Package main - runs code to demonstrate package word
package main

import (
	"fmt"
	"myprojects/013-level-13/02/quote"
	"myprojects/013-level-13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
