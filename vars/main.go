package main

import "fmt"

type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    var count int
    var name string
    var active bool

    fmt.Println(count, name, active) 

    today := Monday
    fmt.Println("Monday is day number:", today) 
}