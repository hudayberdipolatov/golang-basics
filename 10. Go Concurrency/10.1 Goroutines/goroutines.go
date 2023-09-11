package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

func func1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("fun1-->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func func2() {
	for i := 1; i <= 10; i++ {
		fmt.Println("fun2-->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
