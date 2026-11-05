package main

import "fmt"

func fizzBuzz() {
    for i := 1; i <= 30; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}

func searchFirst() {
	nums := []int{67, 43, 86, 897, 22143, 42142}
	for _, n := range nums {
		if n > 20 {
			fmt.Println("Найдено:", n) 
			break
		}
	}
}

func main() {
	fizzBuzz()
	searchFirst()
}