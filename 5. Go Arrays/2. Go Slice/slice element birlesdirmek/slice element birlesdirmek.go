package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	newSlice := []int{}
	newSlice = append(slice1, slice2...)
	fmt.Println("slice1 ------->", slice1)
	fmt.Println("----------------------------")
	fmt.Println("slice1 ------->", slice2)
	fmt.Println("----------------------------")
	fmt.Println("newSlice ----->", newSlice)
}
