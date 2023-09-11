package main

import "fmt"

func main() {
	var a int
	fmt.Print("Yasynyzy girizin :")
	fmt.Scanln(&a)
	if a == 18 {
		fmt.Println("siz 18 yasynyzda")
	} else if a > 18 {
		fmt.Println("siz 18 yasdan uly ! Sizin yasynyz : ", a)
	} else {
		fmt.Println("siz 18 yasdan kici ! Sizin yasynyz : ", a)
	}

}
