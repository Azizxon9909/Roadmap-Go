package main

import (
	"errors"
	"fmt"
	"math"
)

func main1() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	// swap(a, b)
	swapByRef(&a, &b)
	fmt.Println(a, b)

	//====================
	// result := sum(1, 2)
	// result, err := sqrt(-81)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
}

func swap(x, y int) int {
	var o int
	o = x
	x = y
	y = o
	return o
}

func swapByRef(x, y *int) int {
	var o int
	o = *x
	*x = *y
	*y = o
	return o
}

func sum(x int, y int) int {
	return x + y
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Maxfiy son uchun emas")
	}
	return math.Sqrt(x), nil
}
