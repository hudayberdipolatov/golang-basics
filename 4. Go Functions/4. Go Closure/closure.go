package main

import "fmt"

func greet(name string) {
	displayData := func() {
		fmt.Println("Hi " + name)
	}
	displayData()
}
func main() {
	greet("Hudayberdi")
}
