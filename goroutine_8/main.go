package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 让出 CPU 给其他 goroutine
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	//runtime.GOMAXPROCS(2) // 最多使用 2 个核

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
}
