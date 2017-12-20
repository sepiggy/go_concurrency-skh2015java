package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	go loop() // Starts a goroutine
	loop()
	time.Sleep(time.Second) // main func wait for goroutine
}
