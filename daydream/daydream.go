package main

import (
	"fmt"
		"unicode/utf8"
)

func ReverseStr(str string) string {
	runes := []rune(str)
	for i,j := 0,len(runes)-1; i < j; i,j = i+1,j-1 {
		runes[i],runes[j] = runes[j],runes[i]
	}
	return string(runes)
}

func daydream(input string) {
	var (
		s string = input
		pattern[4] string =[4]string{"dream","dreamer","erase","eraser"} 
		rArr[4] string
	)
	r := ReverseStr(s)
	for i := 0; i < len(pattern); i++ {
		revRune := ReverseStr(pattern[i])
		rArr[i] = revRune
	}
	can := true
	for i := 0; i < utf8.RuneCountInString(r); {
		canD := false
		for j := 0; j < len(rArr); j++ {
			d := rArr[j]
			if r[i:utf8.RuneCountInString(d) + i] == d {
				canD = true
				i += utf8.RuneCountInString(d)
				if (i == utf8.RuneCountInString(r)) {
					break
				}
			}
		}
		if (!canD) {
			can = false
			break
		}
	}
	if (can) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	daydream("dreameraser")
	daydream("dreamerer")
}