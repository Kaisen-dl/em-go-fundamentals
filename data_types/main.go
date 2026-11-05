package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var i int = 42
    var f float64 = 4.2
    var s string = "hi"
    var b bool = true

    fmt.Println("int:", i, "size:", unsafe.Sizeof(i), "bytes")
    fmt.Println("float64:", f, "size:", unsafe.Sizeof(f), "bytes")
    fmt.Println("string:", s, "bytes:", len(s), "runes:", len([]rune(s)))
    fmt.Println("bool:", b, "size:", unsafe.Sizeof(b), "bytes")
}