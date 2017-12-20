package main

import "fmt"

var ch chan int = make(chan int)

func foo() {
	ch <- 0
	fmt.Printf("foo function completed\n")
}

func main() {
	go foo()
	// message := <-ch
	// fmt.Println(message)
}
