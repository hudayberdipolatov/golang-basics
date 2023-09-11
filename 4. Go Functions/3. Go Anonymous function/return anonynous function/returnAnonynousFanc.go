package main

import "fmt"

func test(a, b int) func() int {
	//a := 15
	//b := 10
	return func() int {
		c := a + b
		return c
	}
}
func main() {
	a := test(15, 10)
	fmt.Println(a())
}
