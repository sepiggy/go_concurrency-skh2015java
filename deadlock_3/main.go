package main

/*
 * Go 启动的所有 goroutine 里的非缓冲信道一定要一个线里存数据，一个线里取数据，要成对才行
 */

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 1		// c 通道的数据没有被其他 goroutine 读取走，堵塞当前 goroutine
		quit <- 0	// quit 始终没有办法写入数据
	}()

	<- quit // // quit 等待数据的写
}
