package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

type counterFunc func() int

func newCounter() (counterFunc, func()) {
    count := 0
    next := func() int {
        count++
        return count
    }
    reset := func() {
        count = 0
    }
    return next, reset
}

func main() {
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    fmt.Println(double(5))  
    fmt.Println(triple(5))  

    c, reset := newCounter()
    fmt.Println(c()) 
    fmt.Println(c()) 
    reset()
    fmt.Println(c()) 
}