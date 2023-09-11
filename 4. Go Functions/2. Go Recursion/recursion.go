package main

import "fmt"

func rekursiya(number int) {

	if number > 1 {
		fmt.Println(number)
		rekursiya(number - 1)
	} else {
		fmt.Println("rekursiya stop")
	}
}
func main() {
	rekursiya(10)
}
