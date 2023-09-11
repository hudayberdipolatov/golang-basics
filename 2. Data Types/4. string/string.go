package main

import "fmt"

func main() {
	var str1 string
	str1 = "Hello"
	str2 := "How are you today"
	fmt.Println("str1", str1)
	fmt.Println("str2", str2)
	fmt.Println("str1 len():", len(str1))
	fmt.Println("add two string    : ", str1+" "+str2)
}
