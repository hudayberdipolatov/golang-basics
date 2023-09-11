package main

import "fmt"

func main() {
	var age int
Loop:

	fmt.Print("Sizde saylaw hukugy barlygyny yada yoklugyny barlamak ucin yasynyzy girizin : ")
	fmt.Scanln(&age)
	if age > 17 {
		fmt.Println("Sizde saylaw hukugy bar !")
	} else {
		fmt.Println("Sizde saylaw hukugy yok sizin yasynyz 18-den kici !")
		goto Loop
	}

}
