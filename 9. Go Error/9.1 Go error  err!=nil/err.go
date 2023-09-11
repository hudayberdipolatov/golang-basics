package main

import (
	"errors"
	"fmt"
	"math"
)

func kwadratKok(number float64) (float64, error) {

	if number < 0 {
		return 0, errors.New("Math: kwadrat kok almak ucin girizilen san 0-dan uly bolmaly")
	}
	kok := math.Sqrt(number)
	return kok, nil
}
func main() {
	var san float64
	fmt.Print("Kok almak isleyan sanynyzy girizin :")
	fmt.Scanln(&san)
	result, err := kwadratKok(san)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
