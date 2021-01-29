package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	FmtTime = "2006-01-02 15:04:05"
)

func main() {
	c1 := make(chan int, 2)
	c2 := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 2; i++ {
			c1 <- i
		}
		close(c1)
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			c2 <- i
		}
		close(c2)
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < 10; i++ {
		select {
		case x, ok := <-c1:
			fmt.Printf("%s : 从c1读取到 x=%v, ok=%v\n", time.Now().Format(FmtTime), x, ok)
			if !ok {
				c1 = nil
			}
			time.Sleep(500 * time.Millisecond)
		case x, ok := <-c2:
			fmt.Printf("%s : 从c2读取到 x=%v, ok=%v\n", time.Now().Format(FmtTime), x, ok)
			if !ok {
				c2 = nil
			}
			time.Sleep(500 * time.Millisecond)
			//default:
			//	fmt.Printf("%s : 没有读取到信息。。。\n", time.Now().Format(FmtTime))
		}
	}
}
