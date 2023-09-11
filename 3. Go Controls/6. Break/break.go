package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("aylaw sanyny giriz :")
	fmt.Scanln(&a)
	fmt.Print("saklanmaly sany giriz :")
	fmt.Scanln(&b)
	for i := 1; i <= a; i++ {
		if i == b {
			break
		}
		fmt.Println(i)
	}
}
