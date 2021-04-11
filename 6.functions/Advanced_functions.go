package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

// ==============================

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {

	test := func(x int) int {
		return x * -1
	}
	test2(test)

	// =======================

	returnFunc("Hello")()
	
	x := returnFunc("Goodbye")
	x()

}
