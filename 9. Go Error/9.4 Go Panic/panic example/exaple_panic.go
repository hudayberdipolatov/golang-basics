package main

import "fmt"

func div(num int) int {
	if num%2 != 0 {
		panic("sany 2 bolenimizde galyndy emele gelyar")
	}
	result := num / 2
	fmt.Println("san 2 bolunyar ")
	return result
}

func main() {
	var san int
	fmt.Print("san giriz : ")
	fmt.Scanln(&san)
	result := div(san)
	fmt.Println("result : ", result)
}
