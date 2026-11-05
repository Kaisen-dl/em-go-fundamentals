package main

import (
    "errors"
    "fmt"
)

func max(x, y int) (int, error) {
	if x > y {
		return x, nil
	} 
	if x < y {
		return y, nil
	} else {
		return 0, errors.New("числа равны")
	}
}

func filter(nums []int, fn func(int) bool) []int {
    var result []int
    for _, n := range nums {
        if fn(n) {
            result = append(result, n)
        }
    }
    return result
}

func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("деление на ноль")
    }
    return x / y, nil
}

func main() {
    fmt.Println(max(5, 55)) 

    delNa3 := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
        return n%3 == 0
    })
    fmt.Println(delNa3) 

    _, err := divide(5, 0)
    if err != nil {
        fmt.Println("Ошибка:", err)
    }
}