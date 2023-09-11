package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("aylaw sany giriz : ")
	fmt.Scanln(&a)
	fmt.Print("continue sany  giriz : ")
	fmt.Scanln(&b)
	for i := 1; i <= a; i++ {
		if i == b {
			continue
		}
		fmt.Println(i)
	}
}
