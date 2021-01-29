package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 2; i++ {
			c1 <- i
		}
		close(c1)
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < 10; i++ {
		select {
		case x, ok := <-c1:
			fmt.Printf("%s : 从c1读取到 x=%v, ok=%v\n", time.Now().
				Format("2006-01-02 15:04:05"), x, ok)
			time.Sleep(500 * time.Millisecond)
			//default:
			//	fmt.Printf("%s : 没有读取到信息。。。\n", time.Now().
			//		Format("2006-01-02 15:04:05"))
		}
	}
}
