package main

import "fmt"

func main() {
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("Ping!")

	fmt.Println(<-messages) // 取消息
}
