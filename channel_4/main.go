package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // 被关闭的信道会禁止数据流入, 是只读的

	for v := range ch {
		fmt.Println(v)
	}
}
