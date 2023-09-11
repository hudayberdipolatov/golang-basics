package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int

func func1(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go func1("foo: ")
	go func1("bar: ")
	wg.Wait()
	fmt.Println("last count value ", count)
}
