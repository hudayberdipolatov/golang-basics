package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	ini, beyikligi float64
}

func (r Rectangle) area() float64 {
	return r.ini * r.beyikligi
}
func calc(s Shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	rect := Rectangle{12, 15}
	calc(rect)
}
