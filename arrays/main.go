package main

import "fmt"

func avg(avg_arr [4]float64) float64 {
	var s float64
	for _, v := range avg_arr {
		s += v
	}
	return s / float64(len(avg_arr))
}

func main() {
	var s int
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		s += v
	}
	
	fmt.Println("s =", s)

	avg_arr := [4]float64{4.5, 3.5, 1.4, 6.3}
	fmt.Println("avg_arr =", avg(avg_arr))
}