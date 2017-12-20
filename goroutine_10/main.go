package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", id)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(2)

	for i := 0; i < 3; i++ {
		go loop(i)
	}

	for i := 0; i < 3; i++ {
		<-quit
	}
}
