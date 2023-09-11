package main

import (
	"fmt"
)

func main() {
	var cars = []string{"toyota", "lexus", "bmw"}
	fmt.Println(cars)
	fmt.Println("-----------------------------")
	cars = append(cars, "mers", "opel", "VW")
	fmt.Println("append function ulanyp slice element gosmak")
	fmt.Println("-----------------------------")
	fmt.Println(cars)

}
