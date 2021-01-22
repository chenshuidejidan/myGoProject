package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				break
			}
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
