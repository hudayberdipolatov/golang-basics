package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Create")
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	file.WriteString("Hi How are you Today !!!")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	readString := string(stream)
	fmt.Println("Read file.txt string ")
	fmt.Println(readString)
}
