package main

import "fmt"

func main() {
	var cars = [5]string{"toyota", "bmw", "lexus", "mers", "opel"}
	for _, car := range cars {
		fmt.Println("car --->", car)
	}
}
