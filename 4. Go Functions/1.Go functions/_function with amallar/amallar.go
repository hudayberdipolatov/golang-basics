package main

import "fmt"

func gosmak() {
	var x, y int
	fmt.Print("x sany giriz : ")
	fmt.Scanln(&x)
	fmt.Print("y sany giriz : ")
	fmt.Scanln(&y)
	c := x + y
	fmt.Println("c = ", c)
}
func ayyrmak() {
	var x, y int
	fmt.Print("x sany giriz : ")
	fmt.Scanln(&x)
	fmt.Print("y sany giriz : ")
	fmt.Scanln(&y)
	c := x - y
	fmt.Println("c = ", c)
}
func kopeltmek() {
	var x, y int
	fmt.Print("x sany giriz : ")
	fmt.Scanln(&x)
	fmt.Print("y sany giriz : ")
	fmt.Scanln(&y)
	c := x * y
	fmt.Println("c = ", c)
}
func bolmek() {
	var x, y int
	fmt.Print("x sany giriz : ")
	fmt.Scanln(&x)
	fmt.Print("y sany giriz : ")
	fmt.Scanln(&y)
	c := x / y
	fmt.Println("c = ", c)
}

func main() {
Amal:
	var amal int
	fmt.Println("Amal saylan : \n 1. Gosmak \n 2. Ayyrmak \n 3. Kopetmek \n 4. Bolmek \n")
	fmt.Print("amallardan birisinin tertip beligisini giriz : ")
	fmt.Scanln(&amal)
	switch amal {
	case 1:
		gosmak()
	case 2:
		ayyrmak()
	case 3:
		kopeltmek()
	case 4:
		bolmek()
	default:
		fmt.Println("gorkezilen sanlardan birini sayla")
		goto Amal
	}
}
