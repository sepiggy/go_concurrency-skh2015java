package main

import (
	"fmt"
	"time"
)

/*
 * 如果当前 goroutine 不发生阻塞，它是不会让出 CPU 给其他 goroutine 的;
 * 否则会让出 CPU 给其他 goroutine
 */

var quit chan int

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second) // 停顿1s
	quit <- 0               // 发消息：我执行完啦！
}

func main() {
	count := 1000
	quit = make(chan int, count) // 1000 size of buffered channel

	for i := 0; i < count; i++ { // starts 1000 goroutines for running foo function
		go foo(i)
	}

	for i := 0; i < count; i++ { // await for 1000 goroutines finish, then main over
		<-quit
	}
}
