package main

import "fmt"

const pi = 3.14 // global const
func main() {
	const Username = "User" //local const
	fmt.Println(pi)
	fmt.Println(Username)
	test1()
}

func test1() {
	fmt.Println("tazeden cagyrmak pi : ", pi)
}
