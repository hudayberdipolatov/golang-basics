package main

import (
	"fmt"
	"reflect"
)

// is wagtyndaky gornuslerin we uytgeyanleri denemek ucin ulanylyar

func main() {
	age := 23
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))
}
