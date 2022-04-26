package main

import (
	"fmt"
)

var helloWorld string = "hello world from Go!"
var ptrHelloWorld *string = &helloWorld

func main() {
	fmt.Println(helloWorld)
	fmt.Println(ptrHelloWorld)
	fmt.Println(*ptrHelloWorld)
}