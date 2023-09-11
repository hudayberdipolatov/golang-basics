package main

import "fmt"

// defer beyleki funckiyalar yerine yetyanca defer goylan funkiya yerine yetmeyar we ol in sonunda yerine yetirilyar
func main() {
	one := "one"
	two := "two"
	three := "three"
	fmt.Println(one)
	defer fmt.Println(two)
	fmt.Println(three)

}
