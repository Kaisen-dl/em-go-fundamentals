package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func double(x *int) {
	if x == nil {
		return
	}
	*x *= 2
}

func main() {
	x, y := 5, 100
	fmt.Println("before swap:", x, y)
	swap(&x, &y)
	fmt.Println("after swap:", x, y)

	double(&x)
	fmt.Println("after doubling x:", x)
}