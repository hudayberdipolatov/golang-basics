package main

import "fmt"

func main() {
	var n int
	fmt.Print("san giriz : ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

}
