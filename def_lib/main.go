package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
)

func main() {
    fmt.Print("enter username: ")
    scan := bufio.NewScanner(os.Stdin)
    scan.Scan()
    username := scan.Text()
	fmt.Println("username:", username)

    time := time.Now()
    fmt.Print("time now:", time)

	greeting := fmt.Sprintf("Привет, %s! Время: %s\n", username, time)

    os.WriteFile("greeting.txt", []byte(greeting), 0644)
    fmt.Println("Приветствие сохранено в greeting.txt")
}