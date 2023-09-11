package main

import "fmt"

func main() {
	x := 22
	fmt.Println("variable value: ", x)
	fmt.Println("memory address value: ", &x)
	fmt.Println("----------------------------")
	fmt.Println("pointer bilen addressin icindaki elementi almak")
	y := 17
	var z *int
	z = &y
	fmt.Println("z = &y :", z)
	fmt.Println("*z :", *z)
}
