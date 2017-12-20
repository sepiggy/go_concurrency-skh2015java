package main

/*
 * 解决死锁的方法：把没取走的数据取走，没放入的数据放入， 因为无缓冲信道不能承载数据
 */
func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 1
		quit <- 0
	}()

	<- c
	<- quit
}
