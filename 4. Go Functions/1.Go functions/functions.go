package main

import "fmt"

// func 2 number add
func add() {
	var x, y int
	fmt.Print("x sany giriz : ")
	fmt.Scanln(&x)
	fmt.Print("y sany giriz : ")
	fmt.Scanln(&y)
	c := x + y
	fmt.Println("c = ", c)
}
func main() {
	add()
}
