package main

import (
	"fmt"
)

func main() {
	var yakusuu int = 0
	for i := 1; i <= 100; i++ {
		yakusuu = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				yakusuu++
			}
		}
		if yakusuu == 2 {
			fmt.Println(i)
		}
	}
}
