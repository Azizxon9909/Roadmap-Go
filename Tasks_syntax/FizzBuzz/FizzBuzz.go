package main

import (
	"fmt"
)

func main() {
	fb := FB(22)
	fmt.Println(fb)
}

func FB(n int) (string) {
	if n%5 == 0 && n%3 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Fizz"
	} else if n%3 == 0 {
		return "Buzz"
	} else {
		return "NotFizzBuzz"
	}
}
