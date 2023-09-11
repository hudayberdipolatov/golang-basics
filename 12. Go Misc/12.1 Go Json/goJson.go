package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string
	Fullname string
	Car      Car
}

type Car struct {
	Brend string
	Model string
	Year  string
}

//func UserAdd() {
//
//}

func main() {
	var p = User{
		UserName: "Polat",
		Fullname: "Hudayberdi Polatow",
		Car: Car{
			Brend: "lexus",
			Model: "RX350",
			Year:  "2023",
		},
	}
	res, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	fmt.Println("-----------------------------")
	json_unmarshal := `{"UserName":"Polat","Fullname":"Hudayberdi Polatow","Car":{"Brend":"lexus","Model":"RX350","Year":"2023"}}`
	var user User
	json.Unmarshal([]byte(json_unmarshal), &user)
	fmt.Println(user)
}
