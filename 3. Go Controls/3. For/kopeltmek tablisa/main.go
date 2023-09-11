package main

import "fmt"

func main() {
	var a int
	fmt.Println("size nace cenli kopeltmek tablisa gerek ?")
	fmt.Print("gerek bolan sanynyzy giriz : ")
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = ", i, j)
			fmt.Println(i * j)
		}
		fmt.Println("-------------------------------")
	}

}
