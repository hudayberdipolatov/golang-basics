package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}
func main() {
	number := 5
	result := factorial(number)
	fmt.Println("Factorial of", number, "is : ", result)
}
