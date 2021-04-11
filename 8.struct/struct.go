package main

import "fmt"

type car struct {
	Make, Model, Color string
	Year, Weight       int
	Engine             engine
}
type engine struct {
	name string
	hp   int
}

func main() {
	// var myCar car
	myCar := car{Make: "Mersades", Model: "XC99", Color: "White", Year: 2020, Weight: 2342, Engine: engine{"T55", 400}}
	fmt.Println((myCar))
}