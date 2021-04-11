package main

import "fmt"

func changeValue(str *string)  {
	*str ="changed!"
}

func changeValue2(str string)  {
	str ="changed!"
}
func main() {
	toChange :="hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
}
