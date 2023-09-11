package main

import "fmt"

func main() {
	ages := map[string]int{"Agageldi": 10, "Rahat": 8, "Wepa": 6}
	fmt.Println(ages)
	fmt.Println("---------------------------")
	fmt.Println("Add element map")
	ages["Hudayberdi"] = 23
	fmt.Println(ages)
	fmt.Println("---------------------------")
	fmt.Println("get element with for range")

	for index, age := range ages {
		fmt.Println(index, "-->", age)
	}
}
