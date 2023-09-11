package main

import (
	"fmt"
	"reflect"
)

// rune go-da ASCII bahasyny sanlaryn usti bn upjun eder.
// 'A' ==> 65
func main() {
	rune := 'A'
	fmt.Print("ASCII---> ", rune, "\n")
	fmt.Println("-------------")
	fmt.Println(reflect.TypeOf(rune))
}
