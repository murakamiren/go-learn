package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 0; i <= 30; i++ {
		// fmt.Println(i)
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
