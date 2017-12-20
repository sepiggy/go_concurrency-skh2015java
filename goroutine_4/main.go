package main

/*
 * 无缓冲的信道在取消息和存消息的时候都会挂起当前的 goroutine，除非另一端已经准备好
 */

import "fmt"

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	complete <- 0 // 执行完毕通过信道发个消息
}

func main() {
	go loop()
	<- complete // 直到 goroutine 跑完，取到消息，main 在此阻塞
}
