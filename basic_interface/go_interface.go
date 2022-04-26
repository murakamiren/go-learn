package main

import (
	"fmt"
)

//インターフェース
//インターフェースとは、型が持つべきメソッドを規約として定めたもの

type Car struct {
	name string
}

type Bike struct {
	name string
}

func (c Car) Accelerate() {
	fmt.Println(c.name + "が加速する")
}

func (c Car) Brake() {
	fmt.Println(c.name + "がブレーキをかける")
}

//インターフェース定義
//この場合は、Accelerate()とBrake()を持っていればどんな値でもok
type Vehicle interface {
	Accelerate()
	Brake()
}

//func定義
func drive(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Brake()
}

func (b Bike) Accelerate() {
	println(b.name + "が加速雨する")
}

func (b Bike) Brake() {
	println(b.name + "がブレーキをかける")
}


func main() {
	superCar := Car{"スーパーカー"}
	drive(superCar)

	yzf_r3 := Bike{"YZF-R3"}
	drive(yzf_r3)
}