package main

import "fmt"

type Rectangle struct {
	Width float64
	Height float64
}

func area(r Rectangle) float64 {
	return r.Width * r.Height
}

func perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

type Stack struct {
	data []int
}

func (s *Stack) push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	i := len(s.data) - 1
    value := s.data[i]
    s.data = s.data[:i]
    return value, true
}

func main() {
	rect := Rectangle{Width: 5, Height: 4}
	fmt.Println("Area:", area(rect)) 
	fmt.Println("Perim:", perimeter(rect)) 

	stack := Stack{}
	stack.push(5)
	stack.push(10)
	stack.push(15)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}