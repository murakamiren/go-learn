package main

import "fmt"

func product(a int,b int) {
	var sumResult = a * b
	if sumResult % 2 == 0 {
		fmt.Println("Even")
	} else if sumResult % 2 == 1 {
		fmt.Println("Odd")
	}
}

func main() {
	product(2, 10)
	product(3, 4)
	product(3, 7)
}