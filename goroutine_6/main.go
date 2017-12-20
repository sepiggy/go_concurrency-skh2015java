package main

import "fmt"

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	// 开两个 goroutine 跑函数 loop, loop 函数负责打印 10 个数
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
}
