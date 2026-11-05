package main

import "fmt"

func main() {
    n := 15
    if n%2 == 0 {
        fmt.Println(n, "чётное")
    } else {
        fmt.Println(n, "нечётное")
    }

    const Read = 1 << 0
    const Write = 1 << 1
    perm := Read | Write
    fmt.Printf("Permissions: %b\n", perm)
    fmt.Println("Has write:", perm&Write != 0)

    a, b := 17, 5
    fmt.Println("int div:", a/b)          
    fmt.Println("float div:", float64(a)/float64(b)) 
}