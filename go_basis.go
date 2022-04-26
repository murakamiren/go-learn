package main

import (
	"fmt"
)

// 変数宣言
var helloWorld string = "hello world from Go!"

// &変数名で変数のメモリアドレスを取得できる
// pointer変数
// pointer変数に型をつける *型
var ptrHelloWorld *string = &helloWorld

var i int = 20
var ptrI *int = &i


//構造体(struct)
//構造体の宣言
// var 変数名 struct {
// 	フィールド名 フィールドの型
// }
var person struct {
	name string
	age int
	gender string
	weight int
}

func main() {
	// 出力
	fmt.Println(helloWorld)
	fmt.Println(ptrHelloWorld)

	//pointer関数に*をつけることによって、そのポインタ変数が指しているメモリの内容を参照できる
	fmt.Println(*ptrHelloWorld)

	fmt.Println(i)
	fmt.Println(ptrI)
	fmt.Println(*ptrI)

	// %#v を使えば、その値のGo言語の構文上の表現を知れる
	// \nは改行
	fmt.Printf("%#v\n", person)

	// フィールドに値を入れる
	person.name = "Tom"
	fmt.Println(person.name)
	person.age = 20
	person.gender = "male"
	person.weight = 50
	fmt.Printf("%#v\n", person)

	//type
	type PersonType struct {
		name string
		age int
		gender string
		weight int
	}

	Tom := PersonType {
		name: "Tom",
		age: 20,
		gender: "male",
		weight: 50,
	}

	fmt.Printf("%#v\n", Tom)

	//順番さえ合えばフィールド名は省略可能
	Josh := PersonType{"Josh", 30, "female", 43}
	fmt.Printf("%#v\n", Josh)
}