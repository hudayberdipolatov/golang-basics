package main

import (
	"fmt"
)

func main() {
	var cars = [3]string{"toyota", "lexus", "bmw"}
	fmt.Println(cars)
	fmt.Println("--------------------------")
	fmt.Println("numbers with array")
	var numbers [5]string
	numbers[0] = "1"
	numbers[1] = "2"
	numbers[2] = "3"
	numbers[3] = "4"
	numbers[4] = "5"
	fmt.Println(numbers)
	fmt.Println("------------------------")
	fmt.Println("for bilen array element gosmak")
	var addEelement [10]int
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			addEelement[i] = j
		}
	}
	fmt.Println(addEelement)
	fmt.Println("------------")
	fmt.Println("for range with ")
	for index, value := range addEelement {
		fmt.Println(index, "--->", value)
	}

	fmt.Println("-----------------------")
	fullname := [3]string{"Polatow", "Hudayberdi", "Nurgeldiyewic"}
	fmt.Println(fullname)
	fmt.Println("------------------------")
	for _, name := range fullname {
		fmt.Print(" ", name)
	}
}
