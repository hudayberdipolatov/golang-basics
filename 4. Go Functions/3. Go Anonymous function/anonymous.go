package main

import "fmt"

func main() {
	anonim := func(a, b int) int {
		return a + b
	}
	add := anonim(10, 20)
	fmt.Println("a+b = ", add)

}
