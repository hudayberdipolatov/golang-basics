package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     string
	Edu     struct {
		Where string
		When  string
		job   string
	}
}

func main() {
	var p Person
	p.Name = "Hudayberdi"
	p.Surname = "Polatow"
	p.Age = "23"
	p.Edu.Where = "Magtymguly TDU"
	p.Edu.job = "Maglumat ulgamlary we tehnologiyalary"
	p.Edu.When = "2018-2023"

	fmt.Println("Fullname :", p.Name+" "+p.Surname)
	fmt.Println("Age :", p.Age)
	fmt.Println("EDU : \n University : ", p.Edu.Where, "\n Hunar : ", p.Edu.job, "\n Okan dowri : ", p.Edu.When)
}
