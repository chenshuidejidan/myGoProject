package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	times := make([]struct{}, 10)
	for range times {
		go func() {
			handle()
			ch <- struct{}{}
		}()
	}
	// Wait for goroutines to complete.
	//for range times {
	//	<-ch
	//}
	fmt.Println("ok")
}

func handle() {
	time.Sleep(2 * time.Second)
}
